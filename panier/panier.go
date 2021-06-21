package main

import (
	"log"
	"net"
	"verretech-microservices/database"
	"verretech-microservices/panier/panierpb"

	"google.golang.org/grpc"
)

type server struct {
	db *database.Mongo
	panierpb.UnimplementedServicePanierServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}

	panierServer := &server{
		db: database.NewMongoConnection(),
	}

	err = panierServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	gRPCServer := grpc.NewServer()
	panierpb.RegisterServicePanierServer(gRPCServer, panierServer)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
