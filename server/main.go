package main

import (
	"log"
	"net"

	pb "simple-grpc-example/server/ecommerce"

	"google.golang.org/grpc"
)

const port = ":3000"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterProductInfoServer(server, &Server{})

	log.Printf("gRPC 서버 시작: %s", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("서버 실패: %v", err)
	}
}
