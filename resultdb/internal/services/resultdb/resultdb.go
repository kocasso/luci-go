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

package resultdb

import (
	"context"
	"time"

	"google.golang.org/genproto/googleapis/bytestream"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/artifactcontent"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// resultDBServer implements pb.ResultDBServer.
//
// It does not return gRPC-native errors; use DecoratedResultDB with
// internal.CommonPostlude.
type resultDBServer struct {
	generateArtifactURL func(ctx context.Context, requestHost, artifactName string) (url string, expiration time.Time, err error)
}

// Options is resultdb server configuration.
type Options struct {
	// InsecureSelfURLs is set to true to use http:// (not https://) for URLs
	// pointing back to ResultDB.
	InsecureSelfURLs bool

	// ContentHostnameMap maps a Host header of GetArtifact request to a host name
	// to use for all user-content URLs.
	//
	// Special key "*" indicates a fallback.
	ContentHostnameMap map[string]string

	// ArtifactRBEInstance is the name of the RBE instance to use for artifact
	// storage. Example: "projects/luci-resultdb/instances/artifacts".
	ArtifactRBEInstance string
}

// InitServer initializes a resultdb server.
func InitServer(srv *server.Server, opts Options) error {
	contentServer, err := newArtifactContentServer(srv.Context, opts)
	if err != nil {
		return errors.Annotate(err, "failed to create an artifact content server").Err()
	}

	// Serve all possible content hostnames.
	hosts := stringset.New(len(opts.ContentHostnameMap))
	for _, v := range opts.ContentHostnameMap {
		hosts.Add(v)
	}
	for _, host := range hosts.ToSortedSlice() {
		contentServer.InstallHandlers(srv.VirtualHost(host))
	}

	pb.RegisterResultDBServer(srv.PRPC, &pb.DecoratedResultDB{
		Service: &resultDBServer{
			generateArtifactURL: contentServer.GenerateSignedURL,
		},
		Prelude:  internal.CommonPrelude,
		Postlude: internal.CommonPostlude,
	})

	// Register an empty Recorder server only to make the discovery service
	// list it.
	// The actual traffic will be directed to another deployment, i.e. this
	// binary will never see Recorder RPCs.
	// TODO(nodir): replace this hack with a separate discovery Deployment that
	// dynamically fetches discovery documents from other deployments and
	// returns their union.
	pb.RegisterRecorderServer(srv.PRPC, nil)
	pb.RegisterDeriverServer(srv.PRPC, nil)

	srv.PRPC.AccessControl = prpc.AllowOriginAll
	return nil
}

func newArtifactContentServer(ctx context.Context, opts Options) (*artifactcontent.Server, error) {
	if opts.ArtifactRBEInstance == "" {
		return nil, errors.Reason("opts.ArtifactRBEInstance is required").Err()
	}

	conn, err := artifactcontent.RBEConn(ctx)
	if err != nil {
		return nil, err
	}
	bs := bytestream.NewByteStreamClient(conn)

	contentServer := &artifactcontent.Server{
		InsecureURLs: opts.InsecureSelfURLs,
		HostnameProvider: func(requestHost string) string {
			if host, ok := opts.ContentHostnameMap[requestHost]; ok {
				return host
			}
			return opts.ContentHostnameMap["*"]
		},

		ReadCASBlob: func(ctx context.Context, req *bytestream.ReadRequest) (bytestream.ByteStream_ReadClient, error) {
			return bs.Read(ctx, req)
		},
		RBECASInstanceName: opts.ArtifactRBEInstance,
	}

	if err := contentServer.Init(ctx); err != nil {
		return nil, err
	}
	return contentServer, nil
}
