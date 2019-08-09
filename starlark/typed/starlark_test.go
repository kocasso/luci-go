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

package typed

import (
	"testing"

	"go.starlark.net/resolve"
	"go.starlark.net/starlark"

	"go.chromium.org/luci/starlark/starlarktest"
)

func init() {
	resolve.AllowLambda = true
	resolve.AllowNestedDef = true
}

func TestAllStarlark(t *testing.T) {
	t.Parallel()

	var letters = []string{"T", "K", "L", "M"}
	allocLetter := func() (l string) {
		if len(letters) == 0 {
			return "X"
		}
		l, letters = letters[0], letters[1:]
		return
	}

	starlarktest.RunTests(t, starlarktest.Options{
		TestsDir: "testdata",
		Predeclared: starlark.StringDict{
			// typed_list(cb, list): new typed.List using the callback as converter.
			"typed_list": starlark.NewBuiltin("typed_list", func(th *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
				var cb starlark.Callable
				var l *starlark.List
				if err := starlark.UnpackPositionalArgs("typed_list", args, kwargs, 2, &cb, &l); err != nil {
					return nil, err
				}
				var vals []starlark.Value
				if l != nil {
					vals = make([]starlark.Value, l.Len())
					for i := 0; i < l.Len(); i++ {
						vals[i] = l.Index(i)
					}
				}

				// Cache *callbackConverter so that all converters build from same
				// callback have same address. This is used to test list.extend fast
				// path.
				if th.Local("converters") == nil {
					th.SetLocal("converters", map[starlark.Callable]Converter{})
				}
				converters := th.Local("converters").(map[starlark.Callable]Converter)
				if converters[cb] == nil {
					converters[cb] = &callbackConverter{th, cb, allocLetter()}
				}

				return NewList(converters[cb], vals)
			}),
		},
	})
}

type callbackConverter struct {
	th  *starlark.Thread
	cb  starlark.Callable
	typ string
}

func (c *callbackConverter) Convert(x starlark.Value) (starlark.Value, error) {
	return starlark.Call(c.th, c.cb, starlark.Tuple{x}, nil)
}

func (c *callbackConverter) Type() string {
	return c.typ
}
