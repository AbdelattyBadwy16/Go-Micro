package main

import (
	"context"
	"log"
	"net"

	pb "github.com/abdelatty/grpc-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Println("gRPC received:", req.Text)
	return &pb.MessageResponse{
		Reply: "Hello from gRPC 👋, you said: " + req.Text,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &server{})

	log.Println("gRPC service running on :50051")
	grpcServer.Serve(lis)
}
