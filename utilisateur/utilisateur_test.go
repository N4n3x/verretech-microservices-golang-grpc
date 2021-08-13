package main

import (
	"context"
	"testing"
	"verretech-microservices/database"
	"verretech-microservices/utilisateur/utilisateurpb"
)

func TestGetUtilisateurs(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	// -------------------------------
	prodpb := utilisateurpb.UtilisateursRequest{}
	rep, err := serv.GetUtilisateurs(ctx, &prodpb)
	if err != nil && len(rep.Utilisateur) > 0 {
		t.Errorf("Unable to get Produits : %v", err)
	}
}

func TestGetUtilisateur(t *testing.T) {
	serv := &server{
		db: database.NewMongoConnection(),
	}
	err := serv.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Errorf("Unable to connect to db : %v", err)
	}
	ctx := context.Background()
	// -------------------------------
	prodpb := utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail: "alex.hern@mail.com",
		},
	}
	rep, err := serv.GetUtilisateur(ctx, &prodpb)
	if err != nil {
		t.Errorf("Unable to get Produits : %v, %v", err, rep)
	}
}
