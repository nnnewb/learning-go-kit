package transport

import (
	"context"
	"play/stringsvc"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc stringsvc.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return nil, err
		}

		return &UppercaseResponse{
			V: v,
		}, nil
	}
}

func MakeCountEndpoint(svc stringsvc.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return &CountResponse{
			N: v,
		}, nil
	}
}
