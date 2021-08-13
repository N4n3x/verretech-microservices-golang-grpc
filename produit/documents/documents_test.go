package documents

import (
	"testing"
	"verretech-microservices/database"
)

type server struct {
	db *database.Mongo
}

func testFindOne(t *testing.T) {
	produitServer := &server{
		db: database.NewMongoConnection(),
	}

	err := produitServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		t.Fatalf("Unable to connect to db : %v", err)
	}
	produit := Produit{
		Ref: "A0001",
	}
	err = produit.FindOne(*produitServer.db.Database)
	if err != nil {
		t.Fatalf("Not find Produit : %v", err)
	}
	if produit.Nom != "" {
		t.Fatalf("Invalide Produit")
	}
}
