package main

import (
	"log"
	"net"

	"github.com/ArjunMalhotra07/email"
	pb "github.com/ArjunMalhotra07/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, &email.EmailServiceServer{})
	log.Println("Email service is running on port 50051..")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
