package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"verretech-microservices/database"
	"verretech-microservices/produit/documents"
	"verretech-microservices/produit/produitpb"

	"go.mongodb.org/mongo-driver/bson"
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

func (server *server) GetProduitsByTags(ctx context.Context, req *produitpb.ProduitsByTags) (*produitpb.ProduitsResponse, error) {
	fmt.Printf("%+v\n", req.Tags)
	t := []string{}
	for _, v := range req.Tags {
		t = append(t, v)
	}

	fmt.Printf("%+v\n", t)
	query := bson.M{
		"tags": bson.M{"$all": t},
	}
	p, err := documents.FindByTags(*server.db.Database, query)
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

func (server *server) GetProduitByRef(ctx context.Context, req *produitpb.ProduitByRefRequest) (*produitpb.ProduitResponse, error) {
	var produit documents.Produit
	produit.Ref = req.Ref
	err := produit.FindOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var response produitpb.ProduitResponse
	response.Produit = produit.ToProduitPB()
	return &response, nil
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

func (server *server) UpdateProduits(ctx context.Context, req *produitpb.ProduitsRequest) (*produitpb.BoolResponse, error) {
	var p []*documents.Produit
	for _, e := range req.Produits {
		mongoProduit, _ := documents.FromProduitPB(e)
		p = append(p, mongoProduit)
	}

	_, err := documents.UpdateAll(*server.db.Database, p)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}

	response := &produitpb.BoolResponse{
		State: true,
	}

	return response, nil
}

func (server *server) DeleteProduit(ctx context.Context, req *produitpb.ProduitByRefRequest) (*produitpb.BoolResponse, error) {
	var produit documents.Produit
	produit.Ref = req.Ref
	_, err := produit.Delete(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var response produitpb.BoolResponse
	response.State = true
	return &response, nil
}

func main() {
	fmt.Println("Service Produit => Starting...")
	PRODUIT_PORT := os.Getenv("PRODUIT_PORT")
	if PRODUIT_PORT == "" {
		PRODUIT_PORT = "50051"
		fmt.Println("Service Produit => PRODUIT_PORT variable not found, 50051 used")
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+PRODUIT_PORT)
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}
	fmt.Println("Service Produit => HTTP Server OK")
	produitServer := &server{
		db: database.NewMongoConnection(),
	}

	err = produitServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	fmt.Println("Service Produit => Database connection OK")
	gRPCServer := grpc.NewServer()
	produitpb.RegisterServiceProduitServer(gRPCServer, produitServer)
	fmt.Println("Service Produit => GRPC Server OK")
	fmt.Println("Service Produit => Startup complete, listen on port ", PRODUIT_PORT)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}

}
