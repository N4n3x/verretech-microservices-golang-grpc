package main

import (
	"context"
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

func (server *server) UpdatePanier(ctx context.Context, panier *panierpb.PanierRequest) (*panierpb.PanierResponse, error) {
	return &panierpb.PanierResponse{}, nil
}

func (server *server) GetPanier(ctx context.Context, panier *panierpb.ByUtilisateurRequest) (*panierpb.PanierResponse, error) {
	return &panierpb.PanierResponse{}, nil
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
