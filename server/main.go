package main

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	pb "github.com/alextanhongpin/traefik-grpc/proto"
)

type echoServer struct{}

func (s *echoServer) Echo(ctx context.Context, msg *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Text: msg.Text,
	}, nil
}

func main() {

	//
	// CRED
	//
	BackendCert, _ := ioutil.ReadFile("./backend.cert")
	BackendKey, _ := ioutil.ReadFile("./backend.key")

	// Generate Certificate struct
	cert, err := tls.X509KeyPair(BackendCert, BackendKey)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
	}

	// Create credentials
	creds := credentials.NewServerTLSFromCert(&cert)

	// Use Credentials in gRPC server options
	serverOption := grpc.Creds(creds)

	//
	// SERVER
	//
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	grpcServer := grpc.NewServer(serverOption)
	defer grpcServer.Stop()

	pb.RegisterEchoServiceServer(grpcServer, &echoServer{})
	reflection.Register(grpcServer)
	log.Println("listening to server at port *:50051. press ctrl + c to cancel.")
	log.Fatal(grpcServer.Serve(lis))
}
