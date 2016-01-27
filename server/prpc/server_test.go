// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package prpc

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/luci/luci-go/common/prpc"
	"github.com/luci/luci-go/server/middleware"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

type greeterService struct{}

func (s *greeterService) SayHello(c context.Context, req *HelloRequest) (*HelloReply, error) {
	if req.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Name unspecified")
	}

	return &HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}

type calcService struct{}

func (s *calcService) Multiply(c context.Context, req *MultiplyRequest) (*MultiplyResponse, error) {
	return &MultiplyResponse{
		Z: req.X & req.Y,
	}, nil
}

func TestServer(t *testing.T) {
	t.Parallel()

	Convey("Greeter service", t, func() {
		var server Server

		// auth.Authenticator.Authenticate is not designed to be called in tests.
		server.CustomAuthenticator = true

		RegisterGreeterServer(&server, &greeterService{})

		Convey("Register Calc service", func() {
			RegisterCalcServer(&server, &calcService{})
			So(server.ServiceNames(), ShouldResembleV, []string{
				"prpc.Calc",
				"prpc.Greeter",
			})
		})

		Convey("Handlers", func() {
			c := context.Background()
			r := httprouter.New()
			server.InstallHandlers(r, middleware.TestingBase(c))
			res := httptest.NewRecorder()
			hiMsg := bytes.NewBufferString(`name: "Lucy"`)
			req, err := http.NewRequest("POST", "/prpc/prpc.Greeter/SayHello", hiMsg)
			So(err, ShouldBeNil)
			req.Header.Set("Content-Type", mtPRPCText)

			invalidArgument := strconv.Itoa(int(codes.InvalidArgument))
			unimplemented := strconv.Itoa(int(codes.Unimplemented))

			Convey("Works", func() {
				req.Header.Set("Accept", mtPRPCText)
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusOK)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, "0")
				So(res.Body.String(), ShouldEqual, "message: \"Hello Lucy\"\n")
			})

			Convey("Invalid Accept header", func() {
				req.Header.Set("Accept", "blah")
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusNotAcceptable)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, invalidArgument)
			})

			Convey("Invalid header", func() {
				req.Header.Set("X-Bin", "zzz")
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusBadRequest)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, invalidArgument)
			})

			Convey("Malformed request message", func() {
				hiMsg.WriteString("\nblah")
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusBadRequest)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, invalidArgument)
			})

			Convey("Invalid request message", func() {
				hiMsg.Reset()
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusBadRequest)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, invalidArgument)
				So(res.Body.String(), ShouldEqual, "Name unspecified\n")
			})

			Convey("no such service", func() {
				req.URL.Path = "/prpc/xxx/SayHello"
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusNotImplemented)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, unimplemented)
			})
			Convey("no such method", func() {
				req.URL.Path = "/prpc/prpc.Greeter/xxx"
				r.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusNotImplemented)
				So(res.Header().Get(prpc.HeaderGRPCCode), ShouldEqual, unimplemented)
			})
		})
	})
}
