// Copyright 2019 The LUCI Authors.
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

package recorder

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/tasks"
	"go.chromium.org/luci/resultdb/internal/testutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateFinalizeInvocationRequest(t *testing.T) {
	Convey(`TestValidateFinalizeInvocationRequest`, t, func() {
		Convey(`Valid`, func() {
			err := validateFinalizeInvocationRequest(&pb.FinalizeInvocationRequest{
				Name: "invocations/a",
			})
			So(err, ShouldBeNil)
		})

		Convey(`Invalid name`, func() {
			err := validateFinalizeInvocationRequest(&pb.FinalizeInvocationRequest{
				Name: "x",
			})
			So(err, ShouldErrLike, `name: does not match`)
		})
	})
}

func TestFinalizeInvocation(t *testing.T) {
	Convey(`TestFinalizeInvocation`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		recorder := newTestRecorderServer()

		token, err := generateInvocationToken(ctx, "inv")
		So(err, ShouldBeNil)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(UpdateTokenMetadataKey, token))

		Convey(`finalized failed`, func() {
			testutil.MustApply(ctx,
				testutil.InsertInvocation("inv", pb.Invocation_FINALIZED, map[string]interface{}{
					"Interrupted": true,
				}),
			)
			_, err := recorder.FinalizeInvocation(ctx, &pb.FinalizeInvocationRequest{Name: "invocations/inv"})
			So(err, ShouldHaveAppStatus, codes.FailedPrecondition, `has already been finalized with different interrupted flag`)
		})

		Convey(`Idempotent`, func() {
			testutil.MustApply(ctx, testutil.InsertInvocation("inv", pb.Invocation_ACTIVE, nil))

			inv, err := recorder.FinalizeInvocation(ctx, &pb.FinalizeInvocationRequest{Name: "invocations/inv"})
			So(err, ShouldBeNil)
			So(inv.State, ShouldEqual, pb.Invocation_FINALIZING)

			inv, err = recorder.FinalizeInvocation(ctx, &pb.FinalizeInvocationRequest{Name: "invocations/inv"})
			So(err, ShouldBeNil)
			So(inv.State, ShouldEqual, pb.Invocation_FINALIZING)
		})

		Convey(`Success`, func() {
			testutil.MustApply(ctx, testutil.InsertInvocation("inv", pb.Invocation_ACTIVE, nil))
			inv, err := recorder.FinalizeInvocation(ctx, &pb.FinalizeInvocationRequest{Name: "invocations/inv"})
			So(err, ShouldBeNil)
			So(inv.State, ShouldEqual, pb.Invocation_FINALIZING)

			// Read the invocation from Spanner to confirm it's really FINALIZING.
			inv, err = span.ReadInvocationFull(ctx, span.Client(ctx).Single(), "inv")
			So(err, ShouldBeNil)
			So(inv.State, ShouldEqual, pb.Invocation_FINALIZING)

			var taskInv span.InvocationID
			taskID := "finalize/" + span.InvocationID("inv").RowID()
			testutil.MustReadRow(ctx, "InvocationTasks", tasks.TryFinalizeInvocation.Key(taskID), map[string]interface{}{
				"InvocationId": &taskInv,
			})
			So(taskInv, ShouldEqual, span.InvocationID("inv"))
		})
	})
}
