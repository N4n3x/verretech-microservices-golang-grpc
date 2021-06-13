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

/// Remplacé par GetProduits
/// En attente pour suppression
// func GetAllProduits(w http.ResponseWriter, r *http.Request) {
// 	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Unable to connect to server : %v", err)
// 	}
// 	produitClient := produitpb.NewServiceProduitClient(cc)
// 	b := &produitpb.GetAllProduitsRequest{}
// 	res, err := produitClient.GetAllProduits(context.Background(), b)
// 	if err != nil {
// 		log.Fatalf("Unable to get Products: %v", err)
// 	}
// 	var produits []*documents.Produit

// 	for _, pr := range res.Listproduits.Produits {
// 		d, derr := documents.FromProduitPB(pr)
// 		if derr != nil {
// 			log.Fatalf("Unable to get Products: %v", err)
// 		}
// 		produits = append(produits, d)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(produits)
// }

/// GetProduitByRef
// @return Produit
// @param Ref -> référence d'un produit
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produit)
}

/// GetProduits
// @return []produit
// @param tag (tag=tag1,tag2,tag3...) permet de filtrer les résultats par tag
func GetProduits(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

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
	} else {
		b := &produitpb.GetAllProduitsRequest{}
		res, err := produitClient.GetAllProduits(context.Background(), b)
		if err != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		data = res
	}

	var produits []*documents.Produit

	for _, pr := range data.Listproduits.Produits {
		d, derr := documents.FromProduitPB(pr)
		if derr != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		produits = append(produits, d)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produits)

	// fmt.Fprintf(w, "Tags: %v\n", tags)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.Header.Get("Authorization"))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/produit", GetAllProduits)
	myRouter.HandleFunc("/produit", GetProduits)
	myRouter.HandleFunc("/produit/{ref}", GetProduitByRef)
	myRouter.Use(loggingMiddleware)
	myRouter.Use(authMiddleware)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
