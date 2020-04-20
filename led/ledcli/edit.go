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

package ledcli

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/maruel/subcommands"

	"go.chromium.org/luci/common/flag/stringlistflag"
	"go.chromium.org/luci/common/flag/stringmapflag"
	"go.chromium.org/luci/led/job"
	"go.chromium.org/luci/led/ledcmd"
)

func editCmd(opts cmdBaseOptions) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "edit [options]",
		ShortDesc: "edits the userland of a JobDescription",
		LongDesc: `Allows common manipulations to a JobDescription.

Example:

led get-builder ... |
  led edit -d os=Linux -p something=[100] |
  led launch
`,

		CommandRun: func() subcommands.CommandRun {
			ret := &cmdEdit{}
			ret.initFlags(opts)
			return ret
		},
	}
}

type cmdEdit struct {
	cmdBase

	dimensions     stringlistflag.Flag
	properties     stringmapflag.Value
	propertiesAuto stringmapflag.Value
	recipeName     string
	experimental   string

	recipeIsolate string
	recipeCIPDPkg string
	recipeCIPDVer string

	processedDimensions job.DimensionEditCommands

	swarmingHost string
	taskName     string
}

func (c *cmdEdit) initFlags(opts cmdBaseOptions) {
	c.Flags.Var(&c.dimensions, "d",
		"(repeatable) edit a dimension. "+
			"This takes a parameter of `dimension{=,-=,+=}[value[@expiration_secs]]`. "+
			"Specifying '=[value[@expiration_secs]]' will Reset the dimension to the"+
			" Set of values specified with = (repeating this adds to the Set."+
			" To clear the dimension, specify `dimension=`). "+
			"Specifying '-=value' will Delete the value from the dimension. "+
			"Specifying '+=value[@expiration_secs]' will Add that value to the dimension (expiration). "+
			"Operations are applied as Resets, Deletions, Additions. "+
			"If expiration_secs are omitted, all slices will have the dimension.")

	c.Flags.Var(&c.properties, "p",
		"(repeatable) override a recipe property. This takes a parameter of `property_name=json_value`. "+
			"Providing an empty json_value will remove that property.")

	c.Flags.Var(&c.propertiesAuto, "pa",
		"(repeatable) override a recipe property, using the recipe engine autoconvert rule. "+
			"This takes a parameter of `property_name=json_value_or_string`. If json_value_or_string "+
			"cannot be decoded as JSON, it will be used verbatim as the property value. "+
			"Providing an empty json_value will remove that property.")

	c.Flags.StringVar(&c.recipeName, "r", "",
		"override the `recipe` to run.")

	// These three are used by the 'recipe_engine/led' module to pin the user
	// task across nested led invocations.
	c.Flags.StringVar(&c.recipeIsolate, "rbh", "",
		"override the recipe bundle `hash` (if not using CIPD or git). These should be prepared with"+
			" `recipes.py bundle` from the repo containing your desired recipe and then isolating the"+
			" resulting folder contents. The `led edit-recipe-bundle` subcommand does all this"+
			" automatically.")
	c.Flags.StringVar(&c.recipeCIPDPkg, "rpkg", "",
		"override the recipe CIPD `package` (if not using isolated).")
	c.Flags.StringVar(&c.recipeCIPDVer, "rver", "",
		"override the recipe CIPD `version` (if not using isolated).")

	c.Flags.StringVar(&c.swarmingHost, "S", "",
		"override the swarming `host` to launch the task on (i.e. chromium-swarm.appspot.com).")

	c.Flags.StringVar(&c.taskName, "name", "",
		"set the task name of the led job as it will show on swarming.")

	c.Flags.StringVar(&c.experimental, "exp", "",
		"set to `true` or `false` to change the Build.Input.Experimental value. `led` jobs, "+
			"by default, always start as experimental.")

	c.cmdBase.initFlags(opts)
}

func (c *cmdEdit) positionalRange() (min, max int) { return 0, 0 }
func (c *cmdEdit) jobInput() bool                  { return true }

func (c *cmdEdit) validateFlags(ctx context.Context, _ []string, _ subcommands.Env) (err error) {
	c.processedDimensions, err = job.MakeDimensionEditCommands(c.dimensions)
	if err != nil {
		return err
	}

	return
}

func (c *cmdEdit) execute(ctx context.Context, _ *http.Client, inJob *job.Definition) (out interface{}, err error) {
	err = inJob.Edit(func(je job.Editor) {
		je.EditDimensions(c.processedDimensions)
		if host := c.swarmingHost; host != "" {
			je.SwarmingHostname(c.swarmingHost)
		}
		if c.taskName != "" {
			je.TaskName(c.taskName)
		}
	})
	if err == nil {
		err = inJob.HighLevelEdit(func(je job.HighLevelEditor) {
			je.Properties(c.properties, false)
			je.Properties(c.propertiesAuto, true)
			if c.recipeName != "" {
				je.Properties(map[string]string{"recipe": c.recipeName}, true)
			}
			if c.recipeIsolate != "" || c.recipeCIPDPkg != "" || c.recipeCIPDVer != "" {
				pkg, ver, path := inJob.HighLevelInfo().TaskPayload()
				if c.recipeIsolate == "" {
					if c.recipeCIPDPkg != "" {
						pkg = c.recipeCIPDPkg
					}
					if c.recipeCIPDVer != "" {
						ver = c.recipeCIPDVer
					}
				} else {
					pkg = ""
					ver = ""
					inJob.UserPayload.Digest = c.recipeIsolate
					// We assume that the incoming isolate was generated by led, and
					// therefore has `ledcmd.RecipeDirectory` hard-coded into it.
					//
					// See comment on ledcmd.RecipeDirectory for more info.
					path = ledcmd.RecipeDirectory
				}
				je.TaskPayload(pkg, ver, path)
			}
			if c.experimental != "" {
				je.Experimental(c.experimental == "true")
			}
		})
	}
	return inJob, err
}

func (c *cmdEdit) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	return c.doContextExecute(a, c, args, env)
}