package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"verretech-microservices/commande/commandepb"
	"verretech-microservices/commande/documents"
	"verretech-microservices/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct {
	db *database.Mongo
	commandepb.UnimplementedServiceCommandeServer
}

var COMMANDE_PORT string
var ERP_SERV string

func (server *server) Valid(ctx context.Context, panierReq *commandepb.PanierRequest) (*commandepb.CommandeResponse, error) {
	panier := panierReq.Panier
	c := &documents.Commande{
		Panier: panier,
	}
	commande, err := c.Valided(*server.db.Database, ERP_SERV)
	if err != nil {
		return nil, err
	}
	commandePB := commande.ToCommandePB()
	return &commandepb.CommandeResponse{
		Commande: commandePB,
	}, nil
}

func (server *server) Confirm(ctx context.Context, commandeReq *commandepb.ConfirmRequest) (*commandepb.CommandeResponse, error) {
	commande, err := documents.Confirmed(*server.db.Database, commandeReq.CommandeID, commandeReq.UserID)
	if err != nil {
		return nil, err
	}
	commandePB := commande.ToCommandePB()
	return &commandepb.CommandeResponse{
		Commande: commandePB,
	}, nil
}

func (server *server) Cancel(ctx context.Context, commandeReq *commandepb.CommandeRequest) (*commandepb.CommandeResponse, error) {
	commande, err := documents.FromCommandePB(commandeReq.Commande)
	if err != nil {
		return nil, err
	}
	commande.Canceled(*server.db.Database)
	commandePB := commande.ToCommandePB()
	return &commandepb.CommandeResponse{
		Commande: commandePB,
	}, nil
}

func (server *server) GetUserCommandes(ctx context.Context, req *commandepb.ByUtilisateurRequest) (*commandepb.CommandesResponse, error) {
	///TODO
	idUtilisateur, err := primitive.ObjectIDFromHex(req.UtilisateurID)
	fmt.Println("Service Commande commandes => ", idUtilisateur.Hex())
	if err != nil {
		return nil, err
	}
	commandes, err := documents.FindByUserID(*server.db.Database, ctx, idUtilisateur)
	fmt.Println("Service Commande commandes => ", commandes)
	commandespb := []*commandepb.Commande{}
	for _, commande := range commandes {
		commandespb = append(commandespb, commande.ToCommandePB())
	}
	return &commandepb.CommandesResponse{
		Commandes: commandespb,
	}, nil
}

func main() {
	COMMANDE_PORT = os.Getenv("COMMANDE_PORT")
	if COMMANDE_PORT == "" {
		COMMANDE_PORT = "50054"
	}
	ERP_SERV = os.Getenv("ERP_SERV")
	if ERP_SERV == "" {
		ERP_SERV = "erp:50055"
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+COMMANDE_PORT)
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}
	fmt.Println("Service Commande => HTTP Server OK")
	commandeServer := &server{
		db: database.NewMongoConnection(),
	}
	err = commandeServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	fmt.Println("Service Commande => Database connection OK")
	gRPCServer := grpc.NewServer()
	commandepb.RegisterServiceCommandeServer(gRPCServer, commandeServer)
	fmt.Println("Service Commande => GRPC Server OK")
	fmt.Println("Service Commande => startup complete, listen on port ", COMMANDE_PORT)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
