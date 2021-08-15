package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"verretech-microservices/commande/documents"
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

func (server *server) ValidERP(ctx context.Context, commandeReq *erppb.CommandeRequest) (*erppb.CommandeResponse, error) {
	// récupérer les info utilisateur commande.panier.utilisateurID
	// pour chaque produit
	commande, err := documents.FromCommandePB(commandeReq.Commande)
	if err != nil {
		return nil, err
	}
	valid := "valid"
	invalid := "invalid"
	commande.Statut = valid
	for _, article := range commandeReq.Commande.Panier.Article {
		url := "https://api.airtable.com/v0/appjpwR0Jl093ePaL/Produit?view=Grid%20view&fields%5B%5D=Ref&fields%5B%5D=Qte%20%28from%20Stock%29&fields%5B%5D=Nom%20%28from%20PointRetrait%29%20%28from%20Stock%29&maxRecords=1&filterByFormula=%7BRef%7D%20%3D%20%27" + article.ProduitRef + "%27"
		var bearer = "Bearer " + API_KEY //"keyCKETjZguzbEMJs"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
			return nil, err
		}
		defer resp.Body.Close()
		// body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
			return nil, err
		}
		var temp documents_erp.StockRepErp
		// if err := json.Unmarshal(body, &temp); err != nil {
		if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
			fmt.Println("failed to unmarshal:", err, " =>", resp, " url: ", url)
			return nil, err
		}
		if temp.Records[0].Fields.QteStock[0] < article.Qte {
			commande.Statut = invalid
		}
	}
	// fmt.Println("ValidERP commande => ", commande)
	return &erppb.CommandeResponse{
		Commande: commande.ToCommandePB(),
	}, nil
}

func (server *server) ConfirmERP(ctx context.Context, commandeReq *erppb.CommandeRequest) (*erppb.CommandeResponse, error) {
	url := "https://api.airtable.com/v0/appjpwR0Jl093ePaL/Commande"
	var bearer = "Bearer " + API_KEY //"keyCKETjZguzbEMJs"
	// body
	informations := fmt.Sprintf("%v", commandeReq.Commande)
	values := map[string]string{"id": commandeReq.Commande.ID, "Informations": informations}
	json_data, err := json.Marshal(values)
	if nil != err {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil, err
	}
	var res map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Println("failed to unmarshal:", err, " =>", resp, " url: ", url)
		return nil, err
	}
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
		PRODUIT_SERV = "produit:50051"
		fmt.Println("ERP Connector => PRODUIT_SERV variable not found, produit:50051 used")
	}
	API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		API_KEY = "keyCKETjZguzbEMJs"
		fmt.Println("ERP Connector => API_KEY variable not found, default used")
		// log.Fatal("ERP Connector => API_KEY variable not found")
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
