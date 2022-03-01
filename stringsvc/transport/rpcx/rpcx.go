package rpcx

import (
	"context"
	"play/stringsvc"
	"play/stringsvc/transport"

	"github.com/go-kit/kit/endpoint"
	"github.com/smallnest/rpcx/server"
)

type StringServiceRPCX struct {
	uppercase endpoint.Endpoint
	count     endpoint.Endpoint
}

func (s *StringServiceRPCX) Uppercase(ctx context.Context, req *transport.UppercaseRequest, resp *transport.UppercaseResponse) error {
	ret, err := s.uppercase(ctx, *req)
	if err != nil {
		return err
	}
	*resp = *(ret.(*transport.UppercaseResponse))
	return nil
}

func (s *StringServiceRPCX) Count(ctx context.Context, req *transport.CountRequest, resp *transport.CountResponse) error {
	ret, err := s.count(ctx, *req)
	if err != nil {
		return err
	}
	*resp = *(ret.(*transport.CountResponse))
	return nil
}

func applyMiddlewares(ep endpoint.Endpoint, middlewares ...endpoint.Middleware) endpoint.Endpoint {
	if len(middlewares) > 0 {
		return endpoint.Chain(middlewares[0], middlewares[1:]...)(ep)
	}
	return ep
}

func NewServer(svc stringsvc.StringService, middlewares ...endpoint.Middleware) *server.Server {
	srv := server.NewServer()
	if err := srv.Register(&StringServiceRPCX{
		uppercase: applyMiddlewares(transport.MakeUppercaseEndpoint(svc), middlewares...),
		count:     applyMiddlewares(transport.MakeCountEndpoint(svc), middlewares...),
	}, ""); err != nil {
		panic(err)
	}

	return srv
}
