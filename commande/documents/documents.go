package documents

import (
	"context"
	"verretech-microservices/commande/commandepb"
	"verretech-microservices/panier/panierpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

/// Statut: Validé => Confirmé => Livré
///                => Annulé
///         Invalidé

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
func (commande *Commande) Valided(db mongo.Database) {
	///TODO:
	// connexion ERP pour valider le stock
	// place le statut à Validé ou Invalidé
	// si valide, insert en base

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
