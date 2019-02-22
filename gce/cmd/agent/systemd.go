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

package main

import (
	"context"
	"io/ioutil"
	"os/exec"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
)

// Ensure SystemdStrategy implements PlatformStrategy.
var _ PlatformStrategy = &SystemdStrategy{}

// SystemdStrategy is a Linux-specific PlatformStrategy using systemd.
// Implements PlatformStrategy.
type SystemdStrategy struct {
	LinuxStrategy
}

// systemdCfg is the path to write the Swarming bot systemd config.
const systemdCfg = "/etc/systemd/system/swarming-start-bot.service"

// systemdTmpl is the name of the Swarming bot systemd config template asset.
const systemdTmpl = "swarming-start-bot.service.tmpl"

// systemdSrv is the name of the Swarming bot systemd service.
const systemdSrv = "swarming-start-bot"

// autostart configures the given Swarming bot code to be executed on startup
// for the given user, then starts the Swarming bot process.
// Implements PlatformStrategy.
func (*SystemdStrategy) autostart(c context.Context, path, user string) error {
	subs := map[string]string{
		"BotCode": path,
		"User":    user,
	}
	s, err := substitute(c, string(GetAsset(systemdTmpl)), subs)
	if err != nil {
		return errors.Annotate(err, "failed to prepare template %q", systemdTmpl).Err()
	}

	logging.Infof(c, "installing: %s", systemdCfg)
	// 0644 allows the systemd config to be read by all users.
	// Useful when SSHing to the instance.
	if err := ioutil.WriteFile(systemdCfg, []byte(s), 0644); err != nil {
		return errors.Annotate(err, "failed to write: %s", systemdCfg).Err()
	}

	logging.Infof(c, "starting %q", systemdSrv)
	if err := exec.Command("systemctl", "start", systemdSrv).Run(); err != nil {
		return errors.Annotate(err, "failed to start service %q", systemdSrv).Err()
	}
	return nil
}
