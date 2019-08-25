package main

import (
	"award/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address  = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAwardClient(conn)
	ctx , cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Draw(ctx, &pb.DrawRequest{Name:"diuge"})
	if err != nil {
		log.Fatalf("draw error : %v", err)
	}
	log.Printf("award : %s", r.Message)
}