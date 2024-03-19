package pbclient

import (
	"context"
	"errors"
	"go-sample-grpc-client/internal/pkg/core/adapter/pbclientadapter"
	"log"

	"github.com/michaelchristopher-ui/proto-test/sample/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PBClient struct {
}

func NewPBClient() PBClient {
	return PBClient{}
}

func (p PBClient) Request(req pbclientadapter.RequestRequest) (pbclientadapter.RequestResponse, error) {
	client, err := p.newClient()

	if err != nil {
		return pbclientadapter.RequestResponse{}, err
	}

	res, err := client.ExampleFunc(context.Background(), &pb.ExampleFuncRequest{
		ExampleRequestString: req.RequestString,
	})

	if err != nil {
		return pbclientadapter.RequestResponse{}, err
	}

	if res == nil {
		return pbclientadapter.RequestResponse{}, errors.New("res is nil")
	}

	log.Printf("example response: %v", res)
	return pbclientadapter.RequestResponse{
		ResponseString: res.GetExampleResponseString(),
	}, nil
}

func (p PBClient) newClient() (pb.ExampleClient, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewExampleClient(conn)
	if client == nil {
		return nil, errors.New("example client is nil")
	}
	return client, nil
}
