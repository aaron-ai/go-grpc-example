package main

import (
	"context"
	v2 "go-grpc-example/apache/rocketmq/v2"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v2.NewMessagingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SendMessage(ctx, &v2.SendMessageRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Status: %s", r.Status)
}
