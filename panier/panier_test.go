package main

import (
	"context"
	"testing"
	"verretech-microservices/database"
	"verretech-microservices/panier/panierpb"
)

func TestGetPanier(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err == nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	// -------------------------------
	panier := panierpb.ByUtilisateurRequest{
		UtilisateurID: "60ce479caf3a5507179f4e33",
	}
	rep, err := serv.GetPanier(ctx, &panier)
	if err != nil {
		t.Errorf("Unable to get Produits : %v, %v", err, rep)
	}
}
