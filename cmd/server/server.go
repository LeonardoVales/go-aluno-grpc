package main

import (
	"log"
	"net"

	"github.com/LeonardoVales/go-aluno-grpc/pb"
	"github.com/LeonardoVales/go-aluno-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50021")
	if err != nil {
		log.Fatalf("Cound not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}
