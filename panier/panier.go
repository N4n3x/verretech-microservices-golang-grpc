package main

import (
	"context"
	"fmt"
	"log"
	"verretech-microservices/produit/produitpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)

	// p1 := &produitpb.Photo{Url: "http://localhost/images/1"}
	// p2 := &produitpb.Photo{Url: "http://localhost/images/2"}
	// var p []*produitpb.Photo
	// p = append(p, p1, p2)

	// l1 := &localisationpb.Localisation{Adresse: "1 rue du pont", Ville: "ROUEN", Cp: "76000"}
	// l2 := &localisationpb.Localisation{Adresse: "24 rue Victor Hugo", Ville: "ROUEN", Cp: "76000"}
	// pr1 := &pointRetraitpb.PointRetrait{Nom: "PONT 1", Localisation: l1}
	// pr2 := &pointRetraitpb.PointRetrait{Nom: "HUGO 24", Localisation: l2}
	// s1 := &produitpb.Stock{PointRetrait: pr1, Qte: 10}
	// s2 := &produitpb.Stock{PointRetrait: pr2, Qte: 5}
	// var s []*produitpb.Stock
	// s = append(s, s1, s2)

	// produit := &produitpb.Produit{
	// 	ID:          "60ae9c57a04348bbf1d50ded",
	// 	Ref:         "Z99999",
	// 	Description: "Un super produit !",
	// 	Prix:        18.99,
	// 	Photos:      p,
	// 	Stocks:      s,
	// 	Tags:        []string{"mat:verre", "cat:sdb"},
	// }

	// produitRequest := &produitpb.ProduitRequest{Produit: produit}
	b := &produitpb.GetAllProduitsRequest{}
	// res, err := produitClient.UpdateProduit(context.Background(), produitRequest)
	res, err := produitClient.GetAllProduits(context.Background(), b)
	if err != nil {
		log.Fatalf("Unable to create Product: %v", err)
	}

	fmt.Printf("Result: %v\n", res)
}
