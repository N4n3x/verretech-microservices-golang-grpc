package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	documents_erp "verretech-microservices/erp/documents"
	"verretech-microservices/erp/erppb"
	"verretech-microservices/produit/produitpb"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type server struct {
	erppb.UnimplementedServiceERPServer
}

var ERP_API_PORT string
var ERP_GRPC_PORT string
var PRODUIT_SERV string
var API_KEY string

///TODO: Fonction de mise à jour des produits
func UpdateProduits(w http.ResponseWriter, r *http.Request) {
	cc, err := grpc.Dial(PRODUIT_SERV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)

	/// Récupération de la liste de produits dans l'ERP
	url := "https://api.airtable.com/v0/appjpwR0Jl093ePaL/Produit?view=Grid%20view"
	var bearer = "Bearer " + API_KEY //"keyCKETjZguzbEMJs"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	var temp documents_erp.RepErp
	if err := json.Unmarshal(body, &temp); err != nil {
		fmt.Println("failed to unmarshal:", err)
	}
	///

	/// Envoi la liste de produits au Service Produit
	var p []*produitpb.Produit
	for _, r := range temp.Records {
		p = append(p, r.Fields.ToProduitPB())
	}
	pReq := &produitpb.ProduitsRequest{
		Produits: p,
	}
	res, err := produitClient.UpdateProduits(context.Background(), pReq)
	if err != nil {
		log.Println("Unable to update Products: ", err)
	}
	json.NewEncoder(w).Encode(res.State)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/updateProduits", UpdateProduits)
	// fmt.Println("ERP Connector => startup complete, listen on port ", ERP_API_PORT)
	log.Fatal(http.ListenAndServe(":"+ERP_API_PORT, myRouter))
}

func (server *server) ConfirmERP(ctx context.Context, commandeReq *erppb.CommandeRequest) (*erppb.CommandeResponse, error) {
	return &erppb.CommandeResponse{}, nil
}

func (server *server) ValidERP(ctx context.Context, commandeReq *erppb.CommandeRequest) (*erppb.CommandeResponse, error) {
	return &erppb.CommandeResponse{}, nil
}

func main() {
	fmt.Println("ERP Connector => Starting...")
	ERP_API_PORT = os.Getenv("ERP_API_PORT")
	if ERP_API_PORT == "" {
		ERP_API_PORT = "50050"
		fmt.Println("ERP Connector => ERP_API_PORT variable not found, 50050 used")
	}
	ERP_GRPC_PORT = os.Getenv("ERP_GRPC_PORT")
	if ERP_GRPC_PORT == "" {
		ERP_GRPC_PORT = "50055"
		fmt.Println("ERP Connector => ERP_GRPC_PORT variable not found, 50055 used")
	}
	PRODUIT_SERV = os.Getenv("PRODUIT_SERV")
	if PRODUIT_SERV == "" {
		PRODUIT_SERV = "localhost:50051"
		fmt.Println("ERP Connector => PRODUIT_SERV variable not found, localhost:50051 used")
	}
	API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		log.Fatal("ERP Connector => API_KEY variable not found")
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+ERP_GRPC_PORT)
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}
	fmt.Println("ERP Connector => HTTP Server OK")
	erpServer := &server{}
	gRPCServer := grpc.NewServer()
	erppb.RegisterServiceERPServer(gRPCServer, erpServer)
	fmt.Println("ERP Connector => GRPC Server OK")
	fmt.Println("ERP Connector => Startup complete, listen on port ", ERP_API_PORT, " and ", ERP_GRPC_PORT)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
	handleRequests()
}
