package main

import (
	"log"
	"net"

	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port      = ":50051"
	jsonPath  = "./product.json" // Lokasi file JSON
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productService := &services.ProductService{FilePath: jsonPath}
	productPb.RegisterProductServiceServer(grpcServer, productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
