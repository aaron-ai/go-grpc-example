package main

import (
	"context"
	gen "go-grpc-example/foo/bar"
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

	c := gen.NewSayHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Unary(ctx, &gen.SayHelloRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Status: %s", r.ResponseContent)
}
