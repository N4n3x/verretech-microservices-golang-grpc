package main

import (
	"context"
	"fmt"
	"N4n3x/verretech-microservices/produit/produitpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)

	p1 := &produitpb.Photo{Url: "http://localhost/images/1"}
	p2 := &produitpb.Photo{Url: "http://localhost/images/2"}
	var p []*produitpb.Photo
	p = append(p, p1, p2)
	
	l1 := &produitpb.Localisation{Adresse: "1 rue du pont", Ville: "ROUEN", Cp: "76000"}
	l2 := &produitpb.Localisation{Adresse: "24 rue Victor Hugo", Ville: "ROUEN", Cp: "76000"}
	pr1 := &produitpb.PointRetrait{Nom: "PONT 1", Localisation: l1}
	pr2 := &produitpb.PointRetrait{Nom: "HUGO 24", Localisation: l2}
	s1 := &produitpb.Stock{PointRetrait: pr1, Qte: 10}
	s2 := &produitpb.Stock{PointRetrait: pr2, Qte: 5}
	var s []*produitpb.Stock
	s = append(s, s1, s2)

	produit := &produitpb.Produit{
		ID:       "",
		Ref: "AB1234",
		Description: "Un super produit !",
		Prix: 10.99,
		Photos: p,
		Stocks: s,
		Tags: []string{"mat:verre", "cat:sdb"},
	}

	produitRequest :=  &produitpb.ProduitRequest{Produit: produit}
	
	res, err := produitClient.AddProduit(context.Background(), produitRequest)
	if err != nil {
		log.Fatalf("Unable to create Product: %v", err)
	}

	fmt.Printf("Created product: %v\n", res)
}