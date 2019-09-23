// Code generated by svcdec; DO NOT EDIT.

package helloworld

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedGreeter struct {
	// Service is the service to decorate.
	Service GreeterServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedGreeter) SayHello(c context.Context, req *HelloRequest) (rsp *HelloReply, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "SayHello", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.SayHello(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "SayHello", rsp, err)
	}
	return
}
