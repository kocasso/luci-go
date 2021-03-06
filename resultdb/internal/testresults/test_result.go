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

package testresults

import (
	"context"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/resultdb/internal/invocations"
	"go.chromium.org/luci/resultdb/internal/span"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// MustParseName retrieves the invocation ID, unescaped test id, and
// result ID.
//
// Panics if the name is invalid. Should be used only with trusted data.
//
// MustParseName is faster than pbutil.ParseTestResultName.
func MustParseName(name string) (invID invocations.ID, testID, resultID string) {
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "invocations" || parts[2] != "tests" || parts[4] != "results" {
		panic(errors.Reason("malformed test result name: %q", name).Err())
	}

	invID = invocations.ID(parts[1])
	testID = parts[3]
	resultID = parts[5]

	unescaped, err := url.PathUnescape(testID)
	if err != nil {
		panic(errors.Annotate(err, "malformed test id %q", testID).Err())
	}
	testID = unescaped

	return
}

// Read reads specified TestResult within the transaction.
// If the TestResult does not exist, the returned error is annotated with
// NotFound GRPC code.
func Read(ctx context.Context, txn span.Txn, name string) (*pb.TestResult, error) {
	invID, testID, resultID := MustParseName(name)
	tr := &pb.TestResult{
		Name:     name,
		TestId:   testID,
		ResultId: resultID,
		Expected: true,
	}

	var maybeUnexpected spanner.NullBool
	var micros spanner.NullInt64
	var summaryHTML span.Compressed
	err := span.ReadRow(ctx, txn, "TestResults", invID.Key(testID, resultID), map[string]interface{}{
		"Variant":         &tr.Variant,
		"IsUnexpected":    &maybeUnexpected,
		"Status":          &tr.Status,
		"SummaryHTML":     &summaryHTML,
		"StartTime":       &tr.StartTime,
		"RunDurationUsec": &micros,
		"Tags":            &tr.Tags,
	})
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		return nil, appstatus.Attachf(err, codes.NotFound, "%s not found", name)

	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", name).Err()
	}

	tr.SummaryHtml = string(summaryHTML)
	populateExpectedField(tr, maybeUnexpected)
	populateDurationField(tr, micros)
	return tr, nil
}

func populateDurationField(tr *pb.TestResult, micros spanner.NullInt64) {
	tr.Duration = nil
	if micros.Valid {
		tr.Duration = ptypes.DurationProto(time.Duration(1000 * micros.Int64))
	}
}

func populateExpectedField(tr *pb.TestResult, maybeUnexpected spanner.NullBool) {
	tr.Expected = !maybeUnexpected.Valid || !maybeUnexpected.Bool
}
