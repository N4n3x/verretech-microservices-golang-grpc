package main

import (
	"context"
	"testing"
	"verretech-microservices/database"
	"verretech-microservices/produit/produitpb"
)

func TestGetAllProduits(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	// -------------------------------
	prodpb := produitpb.GetAllProduitsRequest{}
	rep, err := serv.GetAllProduits(ctx, &prodpb)
	if err != nil && len(rep.Listproduits.Produits) > 0 {
		t.Errorf("Unable to get Produits : %v", err)
	}
}

func TestGetProduitsByTags(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	//------------------------------
	prodpb := produitpb.ProduitsByTags{
		Tags: []string{"Cuisine"},
	}
	rep, err := serv.GetProduitsByTags(ctx, &prodpb)
	if err != nil && len(rep.Listproduits.Produits) > 0 {
		t.Errorf("Unable to get Produits : %v", err)
	}
}

func TestGetProduitByRef(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	//------------------------------
	prodpb := produitpb.ProduitByRefRequest{
		Ref: "A0001",
	}
	rep, err := serv.GetProduitByRef(ctx, &prodpb)
	if err != nil {
		t.Errorf("Unable to get Produits : %v , %v", err, rep)
	}
}
