package main

import (
	"context"
	gen "go-grpc-example/foo/bar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	gen.UnimplementedSayHelloServer
}

func (s *server) Unary(ctx context.Context, request *gen.SayHelloRequest) (*gen.SayHelloResponse, error) {
	return nil, nil
}

func (s *server) BidirectionalStream(gen.SayHello_BidirectionalStreamServer) error {
	return nil
}
func (s *server) ServerStream(*gen.SayHelloRequest, gen.SayHello_ServerStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	gen.RegisterSayHelloServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
