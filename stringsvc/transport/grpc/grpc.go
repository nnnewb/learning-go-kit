package transport

import (
	"context"
	"play/stringsvc"
	"play/stringsvc/transport"
	"play/stringsvc/transport/grpc/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type set struct {
	*pb.UnimplementedStringServiceServer
	uppercase grpctransport.Handler
	count     grpctransport.Handler
}

func NewGRPCServer(svc stringsvc.StringService) *set {
	return &set{
		uppercase: grpctransport.NewServer(transport.MakeUppercaseEndpoint(svc), decodeUppercaseRequestGRPC, encodeUppercaseResponseGRPC),
		count:     grpctransport.NewServer(transport.MakeCountEndpoint(svc), decodeCountRequestGRPC, encodeCountResponseGRPC),
	}
}

func (s *set) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, resp, err := s.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UppercaseResponse), nil
}

func (s *set) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CountResponse), nil
}

func decodeUppercaseRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	incoming := request.(*pb.UppercaseRequest)
	return transport.UppercaseRequest{
		S: incoming.S,
	}, nil
}

func encodeUppercaseResponseGRPC(_ context.Context, response interface{}) (interface{}, error) {
	outcoming := response.(*transport.UppercaseResponse)
	return &pb.UppercaseResponse{
		V: outcoming.V,
	}, nil
}

func decodeCountRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	incoming := request.(*pb.CountRequest)
	return transport.CountRequest{
		S: incoming.S,
	}, nil
}

func encodeCountResponseGRPC(_ context.Context, response interface{}) (interface{}, error) {
	outcoming := response.(*transport.CountResponse)
	return &pb.CountResponse{
		N: int32(outcoming.N),
	}, nil
}
