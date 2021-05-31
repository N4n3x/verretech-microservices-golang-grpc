package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"verretech-microservices/produit/documents"
	"verretech-microservices/produit/produitpb"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func GetAllProduits(w http.ResponseWriter, r *http.Request) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	b := &produitpb.GetAllProduitsRequest{}
	res, err := produitClient.GetAllProduits(context.Background(), b)
	if err != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}
	var produits []*documents.Produit

	for _, pr := range res.Listproduits.Produits {
		d, derr := documents.FromProduitPB(pr)
		if derr != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		produits = append(produits, d)
	}
	json.NewEncoder(w).Encode(produits)
}

func GetProduitByRef(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ref"]

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	b := &produitpb.ProduitByRefRequest{
		Ref: key,
	}
	res, err := produitClient.GetProduitByRef(context.Background(), b)
	if err != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}
	produit, perr := documents.FromProduitPB(res.Produit)
	if perr != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}
	json.NewEncoder(w).Encode(produit)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/produit", GetAllProduits)
	myRouter.HandleFunc("/produit/{ref}", GetProduitByRef)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
