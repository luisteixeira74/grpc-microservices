package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/luisteixeira74/grpc-microservices/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	var err error

	// Tenta conectar com retries e delay
	for {
		conn, err = grpc.Dial("service-a:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			log.Println("✅ Connected to service-a")
			break
		}

		log.Printf("❌ Failed to connect to service-a: %v", err)
		log.Println("⏳ Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
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
			log.Fatalf("❌ Error receiving: %v", err)
		}

		log.Printf("📝 Received: %s", res.GetWord())
	}
}
