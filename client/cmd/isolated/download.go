// Copyright 2015 The LUCI Authors.
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
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/maruel/subcommands"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/client/downloader"
	"go.chromium.org/luci/common/data/caching/cache"
	"go.chromium.org/luci/common/data/text/units"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/isolatedclient"
	"go.chromium.org/luci/common/system/signals"
)

func cmdDownload(authOpts auth.Options) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "download <options>...",
		ShortDesc: "downloads a file or a .isolated tree from an isolate server.",
		LongDesc: `Downloads one or multiple files, or a isolated tree from the isolate server.

Files are referenced by their hash`,
		CommandRun: func() subcommands.CommandRun {
			c := downloadRun{}
			c.commonFlags.Init(authOpts)
			// TODO(mknyszek): Add support for downloading individual files.
			c.Flags.StringVar(&c.outputDir, "output-dir", ".", "The directory where files will be downloaded to.")
			c.Flags.StringVar(&c.outputFiles, "output-files", "", "File into which the full list of downloaded files is written to.")
			c.Flags.StringVar(&c.isolated, "isolated", "", "Hash of a .isolated tree to download.")

			c.Flags.StringVar(&c.cacheDir, "cache-dir", "", "Cache directory to store downloaded files.")
			c.Flags.Int64Var(&c.maxSize, "cache-max-size", 0, "Cache is trimmed if the cache gets larger than this value.")
			c.Flags.IntVar(&c.maxItems, "cache-max-items", 0, "Maximum number of items to keep in the cache.")
			c.Flags.Int64Var(&c.minFreeSpace, "cache-min-free-space", 0, "Cache is trimmed if disk free space becomes lower than this value.")
			return &c
		},
	}
}

type downloadRun struct {
	commonFlags
	outputDir   string
	outputFiles string
	isolated    string

	cacheDir     string
	maxSize      int64
	maxItems     int
	minFreeSpace int64
}

func (c *downloadRun) Parse(a subcommands.Application, args []string) error {
	if err := c.commonFlags.Parse(); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.New("position arguments not expected")
	}
	if c.isolated == "" {
		return errors.New("isolated is required")
	}

	if c.cacheDir == "" && (c.maxSize != 0 || c.maxItems != 0 || c.minFreeSpace != 0) {
		return errors.New("cache-dir is necessary when cache-max-size, cache-max-items or cache-min-free-space are specified,")
	}
	return nil
}

func (c *downloadRun) main(a subcommands.Application, args []string) error {
	// Prepare isolated client.
	ctx, cancel := context.WithCancel(c.defaultFlags.MakeLoggingContext(os.Stderr))
	signals.HandleInterrupt(cancel)
	authClient, err := c.createAuthClient(ctx)
	if err != nil {
		return err
	}
	client := isolatedclient.New(nil, authClient, c.isolatedFlags.ServerURL, c.isolatedFlags.Namespace, nil, nil)
	var filesMu sync.Mutex
	var files []string

	var diskCache cache.Cache
	if c.cacheDir != "" {
		diskCache, err = cache.NewDisk(cache.Policies{
			MaxSize:      units.Size(c.maxSize),
			MaxItems:     c.maxItems,
			MinFreeSpace: units.Size(c.minFreeSpace),
		}, c.cacheDir, c.isolatedFlags.Namespace)
		if err != nil {
			return err
		}
		defer diskCache.Close()
	}

	dl := downloader.New(ctx, client, isolated.HexDigest(c.isolated), c.outputDir, &downloader.Options{
		FileCallback: func(name string, _ *isolated.File) {
			filesMu.Lock()
			files = append(files, name)
			filesMu.Unlock()
		},
		Cache: diskCache,
	})
	if err := dl.Wait(); err != nil {
		return errors.Annotate(err, "failed to call FetchIsolated()").Err()
	}
	if c.outputFiles != "" {
		filesData := strings.Join(files, "\n")
		if len(files) > 0 {
			filesData += "\n"
		}

		if err := ioutil.WriteFile(c.outputFiles, []byte(filesData), 0664); err != nil {
			return errors.Annotate(err, "failed to call WriteFile(%s, ...)", c.outputFiles).Err()
		}
	}
	return nil
}

func (c *downloadRun) Run(a subcommands.Application, args []string, _ subcommands.Env) int {
	if err := c.Parse(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: failed to call Parse(%s): %v\n", a.GetName(), args, err)
		return 1
	}
	cl, err := c.defaultFlags.StartTracing()
	if err != nil {
		fmt.Fprintf(a.GetErr(), "%s: failed to call StartTracing(): %v\n", a.GetName(), err)
		return 1
	}
	defer cl.Close()
	if err := c.main(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: failed to call main(%s): %v\n", a.GetName(), args, err)
		return 1
	}
	return 0
}
