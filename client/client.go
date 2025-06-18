package main

import (
	"context"
	"log"
	"time"

	pb "edutrack-go-backend/edutrack-backend/gen/edutrackgrpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)

	time.Sleep(10 * time.Second) //добавил задержку чтобы контейнер не падал сразу
}
