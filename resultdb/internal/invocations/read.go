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

package invocations

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// ReadColumns reads the specified columns from an invocation Spanner row.
// If the invocation does not exist, the returned error is annotated with
// NotFound GRPC code.
// For ptrMap see ReadRow comment in span/util.go.
func ReadColumns(ctx context.Context, txn span.Txn, id ID, ptrMap map[string]interface{}) error {
	if id == "" {
		return errors.Reason("id is unspecified").Err()
	}
	err := span.ReadRow(ctx, txn, "Invocations", id.Key(), ptrMap)
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		return appstatus.Attachf(err, codes.NotFound, "%s not found", id.Name())

	case err != nil:
		return errors.Annotate(err, "failed to fetch %s", id.Name()).Err()

	default:
		return nil
	}
}

func readMulti(ctx context.Context, txn span.Txn, ids IDSet, f func(id ID, inv *pb.Invocation) error) error {
	if len(ids) == 0 {
		return nil
	}

	st := spanner.NewStatement(`
		SELECT
		 i.InvocationId,
		 i.State,
		 i.CreatedBy,
		 i.CreateTime,
		 i.FinalizeTime,
		 i.Deadline,
		 i.Tags,
		 i.BigQueryExports,
		 ARRAY(SELECT IncludedInvocationId FROM IncludedInvocations incl WHERE incl.InvocationID = i.InvocationId),
		 i.ProducerResource,
		 i.Realm,
		FROM Invocations i
		WHERE i.InvocationID IN UNNEST(@invIDs)
	`)
	st.Params = span.ToSpannerMap(map[string]interface{}{
		"invIDs": ids,
	})
	var b span.Buffer
	return span.Query(ctx, txn, st, func(row *spanner.Row) error {
		var id ID
		included := IDSet{}
		var bqExports [][]byte
		inv := &pb.Invocation{}

		var createdBy spanner.NullString
		var producerResource spanner.NullString
		var realm spanner.NullString
		err := b.FromSpanner(row, &id,
			&inv.State,
			&createdBy,
			&inv.CreateTime,
			&inv.FinalizeTime,
			&inv.Deadline,
			&inv.Tags,
			&bqExports,
			&included,
			&producerResource,
			&realm)
		if err != nil {
			return err
		}

		inv.Name = pbutil.InvocationName(string(id))
		inv.IncludedInvocations = included.Names()
		inv.CreatedBy = createdBy.StringVal
		inv.ProducerResource = producerResource.StringVal
		inv.Realm = realm.StringVal

		if len(bqExports) > 0 {
			inv.BigqueryExports = make([]*pb.BigQueryExport, len(bqExports))
			for i, buf := range bqExports {
				inv.BigqueryExports[i] = &pb.BigQueryExport{}
				if err := proto.Unmarshal(buf, inv.BigqueryExports[i]); err != nil {
					return errors.Annotate(err, "%s: failed to unmarshal BigQuery export", inv.Name).Err()
				}
			}
		}

		return f(id, inv)
	})
}

// Read reads one invocation from Spanner.
// If the invocation does not exist, the returned error is annotated with
// NotFound GRPC code.
func Read(ctx context.Context, txn span.Txn, id ID) (*pb.Invocation, error) {
	var ret *pb.Invocation
	err := readMulti(ctx, txn, NewIDSet(id), func(id ID, inv *pb.Invocation) error {
		ret = inv
		return nil
	})

	switch {
	case err != nil:
		return nil, err
	case ret == nil:
		return nil, appstatus.Errorf(codes.NotFound, "%s not found", id.Name())
	default:
		return ret, nil
	}
}

// ReadBatch reads multiple invocations from Spanner.
// If any of them are not found, returns an error.
func ReadBatch(ctx context.Context, txn span.Txn, ids IDSet) (map[ID]*pb.Invocation, error) {
	ret := make(map[ID]*pb.Invocation, len(ids))
	err := readMulti(ctx, txn, ids, func(id ID, inv *pb.Invocation) error {
		if _, ok := ret[id]; ok {
			panic("query is incorrect; it returned duplicated invocation IDs")
		}
		ret[id] = inv
		return nil
	})
	if err != nil {
		return nil, err
	}
	for id := range ids {
		if _, ok := ret[id]; !ok {
			return nil, appstatus.Errorf(codes.NotFound, "%s not found", id.Name())
		}
	}
	return ret, nil
}

// ReadState returns the invocation's state.
func ReadState(ctx context.Context, txn span.Txn, id ID) (pb.Invocation_State, error) {
	var state pb.Invocation_State
	err := ReadColumns(ctx, txn, id, map[string]interface{}{"State": &state})
	return state, err
}

// ReadStateBatch reads the states of multiple invocations.
func ReadStateBatch(ctx context.Context, txn span.Txn, ids IDSet) (map[ID]pb.Invocation_State, error) {
	ret := make(map[ID]pb.Invocation_State)
	err := txn.Read(ctx, "Invocations", ids.Keys(), []string{"InvocationID", "State"}).Do(func(r *spanner.Row) error {
		var id ID
		var s pb.Invocation_State
		if err := span.FromSpanner(r, &id, &s); err != nil {
			return errors.Annotate(err, "failed to fetch %s", ids).Err()
		}
		ret[id] = s
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// ReadRealm returns the invocation's realm.
func ReadRealm(ctx context.Context, txn span.Txn, id ID) (string, error) {
	var realm string
	err := ReadColumns(ctx, txn, id, map[string]interface{}{"Realm": &realm})
	return realm, err
}
