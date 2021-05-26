package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"verretech-microservices/database"
	"verretech-microservices/produit/documents"
	"verretech-microservices/produit/produitpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *database.Mongo
	produitpb.UnimplementedServiceProduitServer
}

func (server *server) AddProduit(ctx context.Context, req *produitpb.ProduitRequest) (*produitpb.ProduitResponse, error) {
	mongoProduit, _ := documents.FromProduitPB(req.Produit)
	// fmt.Printf("Mongo produit %+v\n", mongoProduit)
	oid, err := mongoProduit.InsertOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	mongoProduit.ID = oid
	var response produitpb.ProduitResponse
	response.Produit = mongoProduit.ToProduitPB()
	return &response, nil
}

func (server *server) GetAllProduits(ctx context.Context, req *produitpb.GetAllProduitsRequest) (*produitpb.ProduitsResponse, error) {

	p, err := documents.Find(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var produits []documents.Produit
	if err = p.All(ctx, &produits); err != nil {
		log.Fatal(err)
	}
	var response produitpb.ListProduits
	for _, pr := range produits {
		response.Produits = append(response.Produits, pr.ToProduitPB())
	}
	var final produitpb.ProduitsResponse
	final.Listproduits = &response
	return &final, nil
}

func (server *server) UpdateProduit(ctx context.Context, req *produitpb.ProduitRequest) (*produitpb.ProduitResponse, error) {
	mongoProduit, _ := documents.FromProduitPB(req.Produit)

	_, err := mongoProduit.Update(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var response produitpb.ProduitResponse
	response.Produit = mongoProduit.ToProduitPB()
	return &response, nil
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
