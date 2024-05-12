package main

import (
	"context"
	"log"
	"time"

	desc "github.com/encountea/chat-server/pkg/chat_api_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address  = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connetc to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewChatApiV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &desc.CreateRequest{Usernames: []string {"Joe", "Stanley", "Luke"}})
	if err != nil {
		log.Fatalf("Failed to get by id: %v", err)
	}

	log.Printf(color.RedString("User names:\n"), color.GreenString("%+v", r.GetId()))
}
