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

package job

import (
	"path"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	swarmingpb "go.chromium.org/luci/swarming/proto/api"
)

type bbInfo struct {
	*Buildbucket

	userPayload *swarmingpb.CASTree
}

var _ Info = bbInfo{}

func (b bbInfo) SwarmingHostname() string {
	return b.GetBbagentArgs().GetBuild().GetInfra().GetSwarming().GetHostname()
}

func (b bbInfo) TaskName() string {
	return b.GetName()
}

func (b bbInfo) CurrentIsolated() (*swarmingpb.CASTree, error) {
	return b.userPayload, nil
}

func (b bbInfo) Dimensions() (ExpiringDimensions, error) {
	ldims := logicalDimensions{}
	for _, reqDim := range b.BbagentArgs.Build.Infra.Swarming.TaskDimensions {
		exp := reqDim.Expiration
		if exp == nil {
			exp = b.BbagentArgs.Build.SchedulingTimeout
		}
		ldims.update(reqDim.Key, reqDim.Value, exp)
	}
	return ldims.toExpiringDimensions(), nil
}

func (b bbInfo) CIPDPkgs() (ret CIPDPkgs, err error) {
	ret = CIPDPkgs{}
	ret.fromList(b.CipdPackages)
	return
}

func (b bbInfo) Env() (ret map[string]string, err error) {
	ret = make(map[string]string, len(b.EnvVars))
	for _, pair := range b.EnvVars {
		ret[pair.Key] = pair.Value
	}
	return
}

func (b bbInfo) Priority() int32 {
	return b.GetBbagentArgs().GetBuild().GetInfra().GetSwarming().GetPriority()
}

func (b bbInfo) PrefixPathEnv() (ret []string, err error) {
	for _, keyVals := range b.EnvPrefixes {
		if keyVals.Key == "PATH" {
			ret = make([]string, len(keyVals.Values))
			copy(ret, keyVals.Values)
			break
		}
	}
	return
}

func (b bbInfo) Tags() (ret []string) {
	if len(b.ExtraTags) > 0 {
		ret = make([]string, len(b.ExtraTags))
		copy(ret, b.ExtraTags)
	}
	return
}

func (b bbInfo) Experimental() bool {
	return b.GetBbagentArgs().GetBuild().GetInput().GetExperimental()
}

func (b bbInfo) Properties() (ret map[string]string, err error) {
	panic("implement me")
}

func (b bbInfo) GerritChanges() (ret []*bbpb.GerritChange) {
	panic("implement me")
}

func (b bbInfo) GitilesCommit() (ret *bbpb.GitilesCommit) {
	panic("implement me")
}

func (b bbInfo) TaskPayload() (cipdPkg, cipdVers string, pathInTask string) {
	exe := b.GetBbagentArgs().GetBuild().GetExe()
	cipdPkg = exe.GetCipdPackage()
	cipdVers = exe.GetCipdVersion()
	pathInTask = path.Dir(b.GetBbagentArgs().GetExecutablePath())
	return
}
