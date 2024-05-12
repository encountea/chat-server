package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/encountea/chat-server/pkg/chat_api_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatApiV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Usernames: %v", req.Usernames)

	return &desc.CreateResponse{
		Id: 2,
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatApiV1Server(s, &server{})

	log.Printf("Server is listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
