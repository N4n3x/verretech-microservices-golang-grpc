package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"verretech-microservices/commande/commandepb"
	"verretech-microservices/database"

	"google.golang.org/grpc"
)

type server struct {
	db *database.Mongo
	commandepb.UnimplementedServicePanierServer
}

var COMMANDE_PORT string

func main() {
	COMMANDE_PORT = os.Getenv("COMMANDE_PORT")
	if COMMANDE_PORT == "" {
		COMMANDE_PORT = "50054"
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
