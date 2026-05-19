package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	pb "core.wcloud.io/generated/grpcgen" // Update
	"core.wcloud.io/services"             // Update
)

func main() {
	grpcAddr := flag.String("grpc", "0.0.0.0:10001", "grpc addr")
	httpAddr := flag.String("http", "0.0.0.0:10000", "http addr")
	flag.Parse()
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	log.Info("Starting server...")

	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen, err: %v", err)
	}

	log.Info("Serving gRPC on ", *grpcAddr)
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	s := &services.ServerServices{}
	pb.RegisterServerServiceServer(grpcServer, s)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	log.Info("Serving HTTP on ", *httpAddr)
	httpServer := runtime.NewServeMux()
	err = pb.RegisterServerServiceHandlerServer(context.Background(), httpServer, s)
	if err != nil {
		log.Fatalf("failed to register service, err: %v", err)
	}
	log.Fatal(http.ListenAndServe(*httpAddr, httpServer))
}
