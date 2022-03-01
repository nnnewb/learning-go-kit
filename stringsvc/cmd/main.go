package main

import (
	"net"
	"net/http"
	"os"
	"play/stringsvc"
	"play/stringsvc/middleware"
	"play/stringsvc/transport"
	grpctransport "play/stringsvc/transport/grpc"
	"play/stringsvc/transport/grpc/pb"
	httptransport "play/stringsvc/transport/http"
	rpcxtransport "play/stringsvc/transport/rpcx"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var svc stringsvc.StringService = &stringsvc.StringServiceImpl{}
	svc = &middleware.LoggingMiddleware{
		StringService: svc,
		Logger:        logger,
	}

	grp := run.Group{}

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	grp.Add(func() error {
		grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterStringServiceServer(grpcServer, grpctransport.NewGRPCServer(svc))
		reflection.Register(grpcServer)
		logger.Log("msg", "grpc", "addr", ":5000")
		return grpcServer.Serve(lis)
	}, func(e error) {
		logger.Log("err", e.Error())
		err = lis.Close()
		if err != nil {
			logger.Log("err", err.Error())
		}
	})

	httpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grp.Add(func() error {
		uppercase := transport.MakeUppercaseEndpoint(svc)
		uppercase = transport.LoggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)
		uppercaseHandler := kithttp.NewServer(uppercase, httptransport.DecodeUppercaseRequest, httptransport.EncodeResponse)
		http.Handle("/uppercase", uppercaseHandler)

		count := transport.MakeCountEndpoint(svc)
		count = transport.LoggingMiddleware(log.With(logger, "method", "count"))(count)
		countHandler := kithttp.NewServer(count, httptransport.DecodeCountRequest, httptransport.EncodeResponse)
		http.Handle("/count", countHandler)

		logger.Log("msg", "HTTP", "addr", ":8080")
		return http.Serve(httpListener, nil)
	}, func(e error) {
		logger.Log("err", e.Error())
		err = httpListener.Close()
		if err != nil {
			logger.Log("err", err.Error())
		}

	})

	rpcxListener, err := net.Listen("tcp", ":6000")
	if err != nil {
		panic(err)
	}
	grp.Add(func() error {
		srv := rpcxtransport.NewServer(svc)
		logger.Log("msg", "rpcx", "addr", ":6000")
		return srv.ServeListener("tcp", rpcxListener)
	}, func(e error) {
		logger.Log("err", e.Error())
		err = rpcxListener.Close()
		if err != nil {
			logger.Log("err", err.Error())
		}
	})

	err = grp.Run()
	logger.Log("err", err.Error())
}
