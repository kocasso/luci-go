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

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

func validateGetArtifactRequest(req *pb.GetArtifactRequest) error {
	if err := pbutil.ValidateArtifactName(req.Name); err != nil {
		return errors.Annotate(err, "name").Err()
	}

	return nil
}

// GetArtifact implements pb.ResultDBServer.
func (s *resultDBServer) GetArtifact(ctx context.Context, in *pb.GetArtifactRequest) (*pb.Artifact, error) {
	if err := validateGetArtifactRequest(in); err != nil {
		return nil, appstatus.BadRequest(err)
	}

	txn := span.Client(ctx).ReadOnlyTransaction()
	defer txn.Close()

	art, err := span.ReadArtifact(ctx, txn, in.Name)
	if err != nil {
		return nil, err
	}

	if err := s.populateFetchURLs(ctx, art); err != nil {
		return nil, err
	}

	return art, nil
}

// populateFetchURLs populates FetchUrl and FetchUrlExpiration fields
// of the artifacts.
func (s *resultDBServer) populateFetchURLs(ctx context.Context, artifacts ...*pb.Artifact) error {
	for _, a := range artifacts {
		url, exp, err := s.generateArtifactURL(ctx, a.Name)
		if err != nil {
			return err
		}
		a.FetchUrl = url
		a.FetchUrlExpiration = pbutil.MustTimestampProto(exp)
	}
	return nil
}