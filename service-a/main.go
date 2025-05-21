package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/luisteixeira74/grpc-microservices/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type randomServiceServer struct {
	pb.UnimplementedRandomServiceServer
}

func (s *randomServiceServer) GetRandomStream(req *pb.RandomRequest, stream pb.RandomService_GetRandomStreamServer) error {
	words := []string{"golang", "microservice", "grpc", "docker", "protobuf"}
	rand.Seed(time.Now().UnixNano())

	for {
		word := words[rand.Intn(len(words))]
		if err := stream.Send(&pb.RandomResponse{Word: word}); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRandomServiceServer(grpcServer, &randomServiceServer{})
	reflection.Register(grpcServer)

	log.Println("🔌 gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
