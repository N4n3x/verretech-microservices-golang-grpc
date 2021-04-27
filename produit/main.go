package main

import (
	"context"
	"fmt"
	"N4n3x/verretech-microservices/database"
	"N4n3x/verretech-microservices/produit/documents"
	"N4n3x/verretech-microservices/produit/produitpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *database.Mongo
	produitpb.UnimplementedServiceProduitServer
}

func (server *server) AddProduit(ctx context.Context, req *produitpb.ProduitRequest) (*produitpb.Produit, error) {
	mongoProduit, _ := documents.FromProduitPB(req.Produit)
	fmt.Printf("Mongo produit %+v\n",mongoProduit)
	oid, err := mongoProduit.InsertOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	mongoProduit.ID = oid

	return mongoProduit.ToProduitPB(), nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}

	produitServer := &server{
		db: database.NewMongoConnection(),
	}

	err = produitServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	gRPCServer := grpc.NewServer()
	produitpb.RegisterServiceProduitServer(gRPCServer, produitServer)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}