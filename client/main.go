package main

import (
	"context"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/alextanhongpin/traefik-grpc/proto"
)

func main() {
	// Read cert file
	FrontendCert, err := ioutil.ReadFile("./frontend.cert")
	if err != nil {
		log.Fatalf("error reading cert: %s", err.Error())
	}

	// Create CertPool
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(FrontendCert); !ok {
		log.Fatal("error appending cert")
	}

	// Create credentials
	credsClient := credentials.NewClientTLSFromCert(roots, "")

	// Dial with specific transport
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(credsClient)}
	// conn, err := grpc.Dial("127.0.0.1:4000", opts...)
	conn, err := grpc.Dial("frontend.local:4443", grpc.WithTransportCredentials(credsClient))
	if err != nil {
		log.Fatalf("fail to dial: %s", err.Error())
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := client.Echo(ctx, &pb.EchoRequest{
		Text: "hello world",
	})
	if err != nil {
		log.Fatalf("error echo: %s", err.Error())
	}
	log.Printf("got res: %#v", res)
}
