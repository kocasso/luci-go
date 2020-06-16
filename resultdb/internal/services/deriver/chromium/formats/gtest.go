// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formats

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

var (
	// Prefixes that may be present in the test name and must be stripped before forming the base id.
	prefixes = []string{"MANUAL_", "PRE_"}

	// Java base ids aren't actually GTest but use the same launcher output format.
	javaIDRe = regexp.MustCompile(`^[\w.]+#`)

	// Test base ids look like FooTest.DoesBar: "FooTest" is the suite and "DoesBar" the test name.
	baseIDRE = regexp.MustCompile(`^(\w+)\.(\w+)$`)

	// Type parametrized test examples:
	// - MyInstantiation/FooTest/1.DoesBar
	// - FooTest/1.DoesBar
	// - FooType/MyType.DoesBar
	//
	// In the above examples, "FooTest" is the suite, "DoesBar" the test name, "MyInstantiation" the
	// optional instantiation, "1" the index of the type on which the test has been instantiated, if
	// no string representation for the type has been provided, and "MyType" is the user-provided
	// string representation of the type on which the test has been instantiated.
	typeParamRE = regexp.MustCompile(`^((\w+)/)?(\w+)/(\w+)\.(\w+)$`)

	// Value parametrized tests examples:
	// - MyInstantiation/FooTest.DoesBar/1
	// - FooTest.DoesBar/1
	// - FooTest.DoesBar/TestValue
	//
	// In the above examples, "FooTest" is the suite, "DoesBar" the test name, "MyInstantiation" the
	// optional instantiation, "1" the index of the value on which the test has been instantiated, if
	// no string representation for the value has been provided, and "TestValue" is the user-provided
	// string representation of the value on which the test has been instantiated.
	valueParamRE = regexp.MustCompile(`^((\w+)/)?(\w+)\.(\w+)/(\w+)$`)

	// TODO(chanli@): Remove this after crbug.com/1045846 is fixed.
	// This is a synthetic test created by test launcher, not a real test.
	syntheticTestRE = regexp.MustCompile(`^GoogleTestVerification.Uninstantiated(?:Type)?ParamaterizedTestSuite<\w+>$`)

	syntheticTestTag = errors.BoolTag{
		Key: errors.NewTagKey("synthetic test"),
	}
)

// GTestResults represents the structure as described to be generated in
// https://cs.chromium.org/chromium/src/base/test/launcher/test_results_tracker.h?l=83&rcl=96020cfd447cb285acfa1a96c37a67ed22fa2499
// (base::TestResultsTracker::SaveSummaryAsJSON)
//
// Fields not used by Test Results are omitted.
type GTestResults struct {
	AllTests   []string `json:"all_tests"`
	GlobalTags []string `json:"global_tags"`

	// PerIterationData is a vector of run iterations, each mapping test names to a list of test data.
	PerIterationData []map[string][]*GTestRunResult `json:"per_iteration_data"`

	// TestLocations maps test names to their location in code.
	TestLocations map[string]*Location `json:"test_locations"`
}

// GTestRunResult represents the per_iteration_data as described in
// https://cs.chromium.org/chromium/src/base/test/launcher/test_results_tracker.h?l=83&rcl=96020cfd447cb285acfa1a96c37a67ed22fa2499
// (base::TestResultsTracker::SaveSummaryAsJSON)
//
// Fields not used by Test Results are omitted.
type GTestRunResult struct {
	Status        string  `json:"status"`
	ElapsedTimeMs float64 `json:"elapsed_time_ms"`

	LosslessSnippet     bool   `json:"losless_snippet"`
	OutputSnippetBase64 string `json:"output_snippet_base64"`

	// Links are not generated by test_launcher, but harnesses built on top may add them to the json.
	Links map[string]string `json:"links"`
}

// Location describes a code location.
type Location struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

// ConvertFromJSON reads the provided reader into the receiver.
//
// The receiver is cleared and its fields overwritten.
func (r *GTestResults) ConvertFromJSON(ctx context.Context, reader io.Reader) error {
	*r = GTestResults{}
	if err := json.NewDecoder(reader).Decode(r); err != nil {
		return err
	}

	if len(r.AllTests) == 0 {
		return errors.Reason(`missing "all_tests" field in JSON`).Err()
	}

	return nil
}

// ToProtos converts test results in r []*pb.TestResult and updates inv
// in-place accordingly.
// If an error is returned, inv is left unchanged.
//
// Does not populate TestResult.Name or TestResult.ResultId.
func (r *GTestResults) ToProtos(ctx context.Context, testIDPrefix string, inv *pb.Invocation) ([]*TestResult, error) {
	// In theory, we can have multiple iterations. This seems rare in practice, so log if we do see
	// more than one to confirm and track.
	if len(r.PerIterationData) > 1 {
		logging.Infof(ctx, "Got %d GTest iterations", len(r.PerIterationData))
	}

	var ret []*TestResult
	var testNames []string
	buf := &strings.Builder{}
	for _, data := range r.PerIterationData {
		// Sort the test name to make the output deterministic.
		testNames = testNames[:0]
		for name := range data {
			testNames = append(testNames, name)
		}
		sort.Strings(testNames)

		for _, name := range testNames {
			baseName, err := extractGTestParameters(name)
			switch {
			case syntheticTestTag.In(err):
				continue
			case err != nil:
				return nil, errors.Annotate(err,
					"failed to extract test base name and parameters from %q", name).Err()
			}

			testID := testIDPrefix + baseName

			for i, result := range data[name] {
				// Store the processed test result into the correct part of the overall map.
				rpb, err := r.convertTestResult(ctx, testID, name, result, buf)
				if err != nil {
					return nil, errors.Annotate(err,
						"iteration %d of test %s failed to convert run result", i, name).Err()
				}

				// TODO(jchinlee): Verify that it's indeed the case that getting NOTRUN results in the final
				// results indicates the task was incomplete.
				// TODO(jchinlee): Check how unexpected SKIPPED tests should be handled.
				ret = append(ret, &TestResult{TestResult: rpb})
			}
		}
	}

	// The code below does not return errors, so it is safe to make in-place
	// modifications of inv.

	// Populate the tags.
	for _, tag := range r.GlobalTags {
		inv.Tags = append(inv.Tags, pbutil.StringPair("gtest_global_tag", tag))
	}
	inv.Tags = append(inv.Tags, pbutil.StringPair(OriginalFormatTagKey, FormatGTest))

	pbutil.NormalizeInvocation(inv)
	return ret, nil
}

func fromGTestStatus(s string) (status pb.TestStatus, expected bool, err error) {
	switch s {
	case "SUCCESS":
		return pb.TestStatus_PASS, true, nil
	case "FAILURE":
		return pb.TestStatus_FAIL, false, nil
	case "FAILURE_ON_EXIT":
		return pb.TestStatus_FAIL, false, nil
	case "TIMEOUT":
		return pb.TestStatus_ABORT, false, nil
	case "CRASH":
		return pb.TestStatus_CRASH, false, nil
	case "SKIPPED":
		return pb.TestStatus_SKIP, true, nil
	case "EXCESSIVE_OUTPUT":
		return pb.TestStatus_FAIL, false, nil
	case "NOTRUN":
		return pb.TestStatus_SKIP, false, nil
	case "UNKNOWN":
		// TODO(jchinlee): Confirm this is reasonable.
		return pb.TestStatus_ABORT, false, nil
	default:
		// This would only happen if the set of possible GTest result statuses change and resultsdb has
		// not been updated to match.
		return pb.TestStatus_STATUS_UNSPECIFIED, false, errors.Reason("unknown GTest status %q", s).Err()
	}
}

// extractGTestParameters extracts parameters from a test id as a mapping with "param/" keys.
func extractGTestParameters(testID string) (baseID string, err error) {
	var suite, name, instantiation, id string

	// If this is a JUnit tests, don't try to extract parameters.
	// TODO: investigate handling parameters for JUnit tests.
	if match := javaIDRe.FindStringSubmatch(testID); match != nil {
		baseID = testID
		return
	}

	// Tests can be only one of type- or value-parametrized, if parametrized at all.
	if match := typeParamRE.FindStringSubmatch(testID); match != nil {
		// Extract type parameter.
		suite = match[3]
		name = match[5]
		instantiation = match[2]
		id = match[4]
	} else if match := valueParamRE.FindStringSubmatch(testID); match != nil {
		// Extract value parameter.
		suite = match[3]
		name = match[4]
		instantiation = match[2]
		id = match[5]
	} else if match := baseIDRE.FindStringSubmatch(testID); match != nil {
		// Otherwise our test id should not be parametrized, so extract the suite
		// and name.
		suite = match[1]
		name = match[2]
	} else if syntheticTestRE.MatchString(testID) {
		// A synthetic test, skip.
		err = errors.Reason("not a real test").Tag(syntheticTestTag).Err()
	} else {
		// Otherwise test id format is unrecognized.
		err = errors.Reason("test id of unknown format").Err()
		return
	}

	// Strip prefixes from test name if necessary.
	for {
		strippedName := name
		for _, prefix := range prefixes {
			strippedName = strings.TrimPrefix(strippedName, prefix)
		}
		if strippedName == name {
			break
		}
		name = strippedName
	}

	switch {
	case id == "":
		baseID = fmt.Sprintf("%s.%s", suite, name)
	case instantiation == "":
		baseID = fmt.Sprintf("%s.%s/%s", suite, name, id)
	default:
		baseID = fmt.Sprintf("%s.%s/%s.%s", suite, name, instantiation, id)
	}

	return
}

func (r *GTestResults) convertTestResult(ctx context.Context, testID, name string, result *GTestRunResult, buf *strings.Builder) (*pb.TestResult, error) {
	status, expected, err := fromGTestStatus(result.Status)
	if err != nil {
		return nil, err
	}

	tr := &pb.TestResult{
		TestId:   testID,
		Expected: expected,
		Status:   status,
		Tags: pbutil.StringPairs(
			// Store the original Gtest test name.
			"gtest_name", name,
			// Store the original GTest status.
			"gtest_status", result.Status,
			// Store the correct output snippet.
			"lossless_snippet", strconv.FormatBool(result.LosslessSnippet),
		),
	}

	// Do not set duration if it is unknown.
	if result.ElapsedTimeMs != 0 {
		tr.Duration = secondsToDuration(result.ElapsedTimeMs / 1000)
	}

	// Write the summary.
	var snippet string
	if result.OutputSnippetBase64 != "" {
		outputBytes, err := base64.StdEncoding.DecodeString(result.OutputSnippetBase64)
		if err != nil {
			// Log the error, but we shouldn't fail to convert an entire invocation just because we can't
			// convert a summary.
			logging.Warningf(ctx, "Failed to convert OutputSnippetBase64 %q", result.OutputSnippetBase64)
		} else {
			snippet = string(outputBytes)
		}
	}
	buf.Reset()
	err = summaryTmpl.ExecuteTemplate(buf, "gtest", map[string]interface{}{
		"snippet": strings.ToValidUTF8(snippet, string(unicode.ReplacementChar)),
		"links":   result.Links,
	})
	if err != nil {
		return nil, err
	}
	tr.SummaryHtml = buf.String()

	// Store the test code location.
	if loc, ok := r.TestLocations[name]; ok {
		tr.Tags = append(tr.Tags,
			pbutil.StringPair("gtest_file", loc.File),
			pbutil.StringPair("gtest_line", strconv.Itoa(loc.Line)),
		)
	}

	return tr, nil
}
