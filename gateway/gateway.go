package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func GetProduitsByTags(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	params := r.URL.Query().Get("tag")
	data := &produitpb.ProduitsResponse{}
	if params != "" {
		tags := strings.Split(params, ",")
		fmt.Printf("%+v\n", tags)
		b := &produitpb.ProduitsByTags{Tags: tags}
		res, err := produitClient.GetProduitsByTags(context.Background(), b)
		if err != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		data = res
		// json.NewEncoder(w).Encode(tags)
	} else {
		b := &produitpb.GetAllProduitsRequest{}
		res, err := produitClient.GetAllProduits(context.Background(), b)
		if err != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		data = res
		// json.NewEncoder(w).Encode("no tag")
	}

	var produits []*documents.Produit

	for _, pr := range data.Listproduits.Produits {
		d, derr := documents.FromProduitPB(pr)
		if derr != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		produits = append(produits, d)
	}
	json.NewEncoder(w).Encode(produits)

	// fmt.Fprintf(w, "Tags: %v\n", tags)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/produit", GetAllProduits)
	myRouter.HandleFunc("/produit/{ref}", GetProduitByRef)
	myRouter.HandleFunc("/produitByTag", GetProduitsByTags)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
