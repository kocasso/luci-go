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

package protoutil

import (
	"testing"

	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateBuilderID(t *testing.T) {
	t.Parallel()

	Convey("ValidateBuilderID", t, func() {
		Convey("nil", func() {
			err := ValidateBuilderID(nil)
			So(err, ShouldErrLike, "project must match")
		})

		Convey("empty", func() {
			b := &pb.BuilderID{}
			err := ValidateBuilderID(b)
			So(err, ShouldErrLike, "project must match")
		})

		Convey("project", func() {
			b := &pb.BuilderID{}
			err := ValidateBuilderID(b)
			So(err, ShouldErrLike, "project must match")
		})

		Convey("bucket", func() {
			Convey("empty", func() {
				b := &pb.BuilderID{
					Project: "project",
				}
				err := ValidateBuilderID(b)
				So(err, ShouldBeNil)
			})

			Convey("invalid", func() {
				b := &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket!",
				}
				err := ValidateBuilderID(b)
				So(err, ShouldErrLike, "bucket must match")
			})

			Convey("v1", func() {
				b := &pb.BuilderID{
					Project: "project",
					Bucket:  "luci.project.bucket",
				}
				err := ValidateBuilderID(b)
				So(err, ShouldErrLike, "invalid use of v1 bucket in v2 API")
			})
		})
	})
}
