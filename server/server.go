package main

import (
	"context"
	v2 "go-grpc-example/apache/rocketmq/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	v2.UnimplementedMessagingServiceServer
}

func (s *server) QueryRoute(ctx context.Context, request *v2.QueryRouteRequest) (*v2.QueryRouteResponse, error) {
	return nil, nil
}

func (s *server) Heartbeat(ctx context.Context, request *v2.HeartbeatRequest) (*v2.HeartbeatResponse, error) {
	return nil, nil

}

func (s *server) SendMessage(ctx context.Context, request *v2.SendMessageRequest) (*v2.SendMessageResponse, error) {
	return nil, nil

}

func (s *server) QueryAssignment(ctx context.Context, request *v2.QueryAssignmentRequest) (*v2.QueryAssignmentResponse, error) {
	return nil, nil

}

func (s *server) ReceiveMessage(request *v2.ReceiveMessageRequest, messageServer v2.MessagingService_ReceiveMessageServer) error {
	return nil
}

func (s *server) AckMessage(ctx context.Context, request *v2.AckMessageRequest) (*v2.AckMessageResponse, error) {
	return nil, nil
}

func (s *server) ForwardMessageToDeadLetterQueue(ctx context.Context, request *v2.ForwardMessageToDeadLetterQueueRequest) (*v2.ForwardMessageToDeadLetterQueueResponse, error) {
	return nil, nil
}

func (s *server) EndTransaction(ctx context.Context, request *v2.EndTransactionRequest) (*v2.EndTransactionResponse, error) {
	return nil, nil
}

func (s *server) Telemetry(telemetryServer v2.MessagingService_TelemetryServer) error {
	return nil
}

func (s *server) NotifyClientTermination(ctx context.Context, request *v2.NotifyClientTerminationRequest) (*v2.NotifyClientTerminationResponse, error) {
	return nil, nil
}

func (s *server) ChangeInvisibleDuration(ctx context.Context, request *v2.ChangeInvisibleDurationRequest) (*v2.ChangeInvisibleDurationResponse, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	v2.RegisterMessagingServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
