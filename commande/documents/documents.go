package documents

import (
	"context"
	"fmt"
	"time"
	"verretech-microservices/commande/commandepb"
	"verretech-microservices/erp/erppb"
	"verretech-microservices/panier/panierpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// Nom de la collection
const commandeCollection = "commande"

// type Article struct {
// 	ProduitRef string `bson:"produitRef"`
// 	Qte        int32  `bson:"quantite"`
// }

// type Panier struct {
// 	ID            *primitive.ObjectID `bson:"_id"`
// 	UtilisateurID primitive.ObjectID  `bson:"utilisateurID"`
// 	Articles      []*Article          `bson:"articles"`
// }

type Commande struct {
	ID        *primitive.ObjectID `bson:"_id"`
	Panier    *panierpb.Panier    `bson:"panier"`
	Timestamp *int64              `bson:"timestamp"`
	Statut    *string             `bson:"statut"`
	Ref       *string             `bson:"ref"`
}

/// Statut: valid => confim => devilery
///                => cancel
///         invalid

// InsertOne Insert une commande en base de données
// Retourne ObjectID de la commande si l'insertion se passe bien, ou une erreur
func (commande *Commande) InsertOne(db mongo.Database) (primitive.ObjectID, error) {
	collection := db.Collection(commandeCollection)
	result, err := collection.InsertOne(context.Background(), commande)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

//@return Commandes if user ID match
func FindByUserID(db mongo.Database, ctx context.Context, id primitive.ObjectID) ([]Commande, error) {
	commandes := []Commande{}
	collection := db.Collection(commandeCollection)
	query := bson.M{"panier": bson.M{"utilisateurID": id}}
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &commandes); err != nil {
		return nil, err
	}
	return commandes, nil
}

//@return Commande with Validé or Invalidé statut
func (commande *Commande) Valided(db mongo.Database, port string) error {
	///TODO:
	// connexion ERP pour valider le stock
	// place le statut à valid ou invalid
	// si valid, insert en base

	cc, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	erpClient := erppb.NewServiceERPClient(cc)

	b := &erppb.CommandeRequest{
		Commande: commande.ToCommandePB(),
	}
	res, err := erpClient.ValidERP(context.Background(), b)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	commande, err = FromCommandePB(res.Commande)
	if err != nil {
		return err
	}
	if res.Commande.Statut == "Valid" {
		now := time.Now()
		ts := now.Unix()
		commande.Timestamp = &ts
		id, err := commande.InsertOne(db)
		if err != nil {
			return err
		}
		commande.ID = &id
	}
	return nil
}

func (commande *Commande) Canceled(db mongo.Database) {
	///TODO:
	// Place le statut Annulé si son statut est à l'état Validé
	// Met à jour en base de données
}

func (commande *Commande) Confirmed(db mongo.Database) {
	///TODO:
	// Place le statut Confirmé si son statut est à l'état Validé
	// Met à jour en base de données
	// connexion ERP pour insertion nouvelle commande
}

func (commande *Commande) ToCommandePB() *commandepb.Commande {
	commandeResp := &commandepb.Commande{}
	if commande.ID != nil {
		commandeResp.ID = commande.ID.Hex()
	}
	if commande.Panier != nil {
		commandeResp.Panier = commande.Panier
	}
	if commande.Ref != nil {
		commandeResp.Ref = *commande.Ref
	}
	if commande.Statut != nil {
		commandeResp.Statut = *commande.Statut
	}
	if commande.Timestamp != nil {
		commandeResp.Timestamp = *commande.Timestamp
	}
	return commandeResp
}

func FromCommandePB(commandeProto *commandepb.Commande) (*Commande, error) {
	commande := &Commande{}
	if commandeProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(commandeProto.ID)
		commande.ID = &oid
	}
	if commandeProto.Panier != nil {
		commande.Panier = commandeProto.Panier
	}
	if commandeProto.Ref != "" {
		commande.Ref = &commandeProto.Ref
	}
	if commandeProto.Statut != "" {
		commande.Statut = &commandeProto.Statut
	}
	if commandeProto.Timestamp != primitive.NilObjectID.Timestamp().Unix() {
		commande.Timestamp = &commandeProto.Timestamp
	}
	return commande, nil
}
