package grpc

import (
	"context"
	"go-sample-grpc-server/pkg/pb"

	"google.golang.org/grpc"
)

type APIIntegrator struct {
	pb.UnimplementedExampleServer
}

func API(s grpc.ServiceRegistrar) {
	pb.RegisterExampleServer(s, NewAPIIntegrator())
}

func NewAPIIntegrator() APIIntegrator {
	return APIIntegrator{}
}

func (a APIIntegrator) ExampleFunc(_ context.Context, req *pb.ExampleFuncRequest) (*pb.ExampleFuncResponse, error) {
	return &pb.ExampleFuncResponse{
		ExampleResponseString: req.GetExampleRequestString() + " as a response",
	}, nil
}
