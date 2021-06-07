package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	documents_erp "verretech-microservices/erp/documents"
	"verretech-microservices/produit/produitpb"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

///TODO: Fonction de mise à jour des produits
func UpdateProduits(w http.ResponseWriter, r *http.Request) {
	log.Println("ok")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)

	/// Récupération de la liste de produits dans l'ERP
	url := "https://api.airtable.com/v0/appjpwR0Jl093ePaL/Produit?view=Grid%20view"
	var bearer = "Bearer " + "keyCKETjZguzbEMJs"
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
		log.Fatalf("Unable to update Products: %v", err)
	}
	json.NewEncoder(w).Encode(res.State)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/updateProduits", UpdateProduits)
	log.Fatal(http.ListenAndServe(":50050", myRouter))
}

func main() {
	handleRequests()
}

// func updateProduct() {
// 	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Unable to connect to server : %v", err)
// 	}
// 	produitClient := produitpb.NewServiceProduitClient(cc)

// 	// p1 := &produitpb.Photo{Url: "http://localhost/images/1"}
// 	// p2 := &produitpb.Photo{Url: "http://localhost/images/2"}
// 	// var p []*produitpb.Photo
// 	// p = append(p, p1, p2)

// 	// l1 := &localisationpb.Localisation{Adresse: "1 rue du pont", Ville: "ROUEN", Cp: "76000"}
// 	// l2 := &localisationpb.Localisation{Adresse: "24 rue Victor Hugo", Ville: "ROUEN", Cp: "76000"}
// 	// pr1 := &pointRetraitpb.PointRetrait{Nom: "PONT 1", Localisation: l1}
// 	// pr2 := &pointRetraitpb.PointRetrait{Nom: "HUGO 24", Localisation: l2}
// 	// s1 := &produitpb.Stock{PointRetrait: pr1, Qte: 10}
// 	// s2 := &produitpb.Stock{PointRetrait: pr2, Qte: 5}
// 	// var s []*produitpb.Stock
// 	// s = append(s, s1, s2)

// 	// produit := &produitpb.Produit{
// 	// 	ID:          "60ae9c57a04348bbf1d50ded",
// 	// 	Ref:         "Z99999",
// 	// 	Description: "Un super produit !",
// 	// 	Prix:        18.99,
// 	// 	Photos:      p,
// 	// 	Stocks:      s,
// 	// 	Tags:        []string{"mat:verre", "cat:sdb"},
// 	// }

// 	// produitRequest := &produitpb.ProduitRequest{Produit: produit}
// 	// b := &produitpb.GetAllProduitsRequest{}
// 	b := &produitpb.ProduitByRefRequest{Ref: "Z99999"}
// 	// res, err := produitClient.UpdateProduit(context.Background(), produitRequest)
// 	// res, err := produitClient.GetAllProduits(context.Background(), b)
// 	// res, err := produitClient.GetProduitByRef(context.Background(), b)
// 	res, err := produitClient.DeleteProduit(context.Background(), b)
// 	if err != nil {
// 		log.Fatalf("Unable to create Product: %v", err)
// 	}

// 	fmt.Printf("Result: %v\n", res)
// }
