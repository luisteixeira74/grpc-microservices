package main

import (
	"context"
	"io"
	"log"
	pb "service-b/proto/randompb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("service-a:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Could not connect to service-a: %v", err)
	}
	defer conn.Close()

	client := pb.NewRandomServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := client.GetRandomStream(ctx, &pb.RandomRequest{})
	if err != nil {
		log.Fatalf("❌ Error calling GetRandomStream: %v", err)
	}

	log.Println("📡 Receiving stream of random words from service-a:")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("✔️ Stream ended")
			break
		}
		if err != nil {
			log.Fatalf("❌ Error receiving stream: %v", err)
		}
		log.Printf("📝 Received: %s", res.GetWord())
	}
}
