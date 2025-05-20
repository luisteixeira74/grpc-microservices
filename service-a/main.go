package main

import (
	"log"
	"math/rand"
	"net"
	pb "service-a/proto/randompb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type randomServiceServer struct {
	pb.UnimplementedRandomServiceServer
}

func (s *randomServiceServer) GetRandomStream(req *pb.RandomRequest, stream pb.RandomService_GetRandomStreamServer) error {
	words := []string{"golang", "microservice", "grpc", "docker", "protobuf"}

	for {
		word := words[rand.Intn(len(words))]
		err := stream.Send(&pb.RandomResponse{Word: word})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRandomServiceServer(s, &randomServiceServer{})
	reflection.Register(s)

	log.Println("🔌 gRPC server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
