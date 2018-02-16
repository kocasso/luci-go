// Copyright 2018 The LUCI Authors.
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

package platform

import (
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCurrentResolution(t *testing.T) {
	t.Parallel()

	Convey("Sanity check on arch/os", t, func() {
		Convey("Has a known os", func() {
			known := false
			switch currentOS {
			case runtime.GOOS, "mac":
				known = true
			}
			So(known, ShouldBeTrue)
		})

		Convey("Has a known architecture", func() {
			known := false
			switch currentArchitecture {
			case runtime.GOARCH, "armv6l":
				known = true
			}
			So(known, ShouldBeTrue)
		})
	})
}
