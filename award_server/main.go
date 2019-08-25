package main

import (
	"award/pb"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)


type server struct {}

const (
	port = ":50051"
)

func (s *server) Draw(ctx context.Context, req *pb.DrawRequest) (*pb.DrawReply, error) {
	name := req.GetName()

	awardBatch := WinPrize(name)
	msg := ""

	if awardBatch == nil {
		msg = "sorry you didn't win any prize"
	} else {
		msg = "congratutions ! you won a %s"
	}

	return &pb.DrawReply{
		Message : msg ,
	} , nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAwardServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//
//}