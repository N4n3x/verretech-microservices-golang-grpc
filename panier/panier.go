package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"verretech-microservices/database"
	"verretech-microservices/panier/documents"
	"verretech-microservices/panier/panierpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct {
	db *database.Mongo
	panierpb.UnimplementedServicePanierServer
}

var PANIER_PORT string

func (server *server) UpdatePanier(ctx context.Context, panierReq *panierpb.PanierRequest) (*panierpb.PanierResponse, error) {
	panier, err := documents.FromPanierPB(panierReq.Panier)
	if err != nil {
		return &panierpb.PanierResponse{}, err
	}
	_, err = panier.Update(*server.db.Database)
	if err != nil {
		return &panierpb.PanierResponse{}, err
	}
	idUtilisateur, err := primitive.ObjectIDFromHex(panierReq.Panier.UtilisateurID)
	var panierRep documents.Panier
	panierRep.FindByID(*server.db.Database, idUtilisateur)
	// var panierInter documents.Panier
	// panierRep.Decode(panierInter)
	// fmt.Printf("Update: %v\n", panierRep)
	return &panierpb.PanierResponse{
		Panier: panierRep.ToPanierPB(),
	}, nil
}

func (server *server) GetPanier(ctx context.Context, req *panierpb.ByUtilisateurRequest) (*panierpb.PanierResponse, error) {
	idUtilisateur, err := primitive.ObjectIDFromHex(req.UtilisateurID)
	if err != nil {
		return nil, err
	}
	var panier documents.Panier
	panier.FindByID(*server.db.Database, idUtilisateur)
	p := panier.ToPanierPB()
	return &panierpb.PanierResponse{
		Panier: p,
	}, nil
}

func main() {
	PANIER_PORT = os.Getenv("PANIER_PORT")
	if PANIER_PORT == "" {
		PANIER_PORT = "50053"
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+PANIER_PORT)
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}
	fmt.Println("Service Panier => HTTP Server OK")
	panierServer := &server{
		db: database.NewMongoConnection(),
	}
	err = panierServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	fmt.Println("Service Panier => Database connection OK")
	gRPCServer := grpc.NewServer()
	panierpb.RegisterServicePanierServer(gRPCServer, panierServer)
	fmt.Println("Service Panier => GRPC Server OK")
	fmt.Println("Service Panier => startup complete, listen on port ", PANIER_PORT)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
