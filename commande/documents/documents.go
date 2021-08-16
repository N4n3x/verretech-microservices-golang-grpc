package documents

import (
	"context"
	"errors"
	"fmt"
	"os"
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
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Panier    *panierpb.Panier   `bson:"panier"`
	Timestamp int64              `bson:"timestamp"`
	Statut    string             `bson:"statut"`
	Ref       string             `bson:"ref"`
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

// UpdateOne Insert une commande en base de données
// Retourne ObjectID de la commande si l'insertion se passe bien, ou une erreur
func (commande *Commande) UpdateOne(db mongo.Database) (int64, error) {
	collection := db.Collection(commandeCollection)
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": commande.ID.Hex()}, commande)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

//@return Commandes if user ID match
func FindByUserID(db mongo.Database, ctx context.Context, id primitive.ObjectID) ([]Commande, error) {
	commandes := []Commande{}
	collection := db.Collection(commandeCollection)
	// query := bson.M{"panier": bson.M{"utilisateurid": id.Hex()}}
	query := bson.M{"panier.utilisateurid": id.Hex()}
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
func (commande *Commande) Valided(db mongo.Database, port string) (Commande, error) {
	///TODO:
	// connexion ERP pour valider le stock
	// place le statut à valid ou invalid
	// si valid, insert en base

	cc, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error grpc: %v", err)
		return Commande{}, err
	}
	erpClient := erppb.NewServiceERPClient(cc)

	b := &erppb.CommandeRequest{
		Commande: commande.ToCommandePB(),
	}
	res, err := erpClient.ValidERP(context.Background(), b)
	if err != nil {
		fmt.Printf("Error validERP: %v", err)
		return Commande{}, err
	}
	fmt.Printf("Valided befor %v", res)
	commandeRep, err := FromCommandePB(res.Commande)
	if err != nil {
		return Commande{}, err
	}
	fmt.Printf("Valided after %v", commande)
	if res.Commande.Statut == "valid" {
		now := time.Now()
		ts := now.Unix()
		commandeRep.Timestamp = ts
		id, err := commandeRep.InsertOne(db)
		if err != nil {
			fmt.Printf("Error insert: %v", err)
			return Commande{}, err
		}
		commandeRep.ID = id
	}
	return *commandeRep, nil
}

func (commande *Commande) Canceled(db mongo.Database) {
	///TODO:
	// Place le statut Annulé si son statut est à l'état Validé
	// Met à jour en base de données
}

func Confirmed(db mongo.Database, idCmd string, idUser string) (Commande, error) {
	///TODO:
	// get commande
	commande := Commande{}
	collection := db.Collection(commandeCollection)
	// query := bson.M{"panier": bson.M{"utilisateurid": id.Hex()}}
	query := bson.M{
		"$and": []bson.M{
			{"panier.utilisateurid": idUser},
			{"_id": idCmd},
		},
	}
	err := collection.FindOne(context.Background(), query).Decode(commande)
	if err != nil {
		return Commande{}, err
	}
	if commande.Statut != "valid" {
		return Commande{}, errors.New("Commande non valide")
	}
	// Place le statut Confirmé si son statut est à l'état Validé
	commande.Statut = "confirm"
	// Met à jour en base de données
	n, err := commande.UpdateOne(db)
	if n <= 0 {
		fmt.Printf("Error grpc: %v", err)
		return Commande{}, errors.New("Aucune commande modifié")
	}
	// connexion ERP pour insertion nouvelle commande
	cc, err := grpc.Dial(os.Getenv("ERP_SERV"), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error grpc: %v", err)
		return Commande{}, err
	}
	erpClient := erppb.NewServiceERPClient(cc)

	b := &erppb.CommandeRequest{
		Commande: commande.ToCommandePB(),
	}
	res, err := erpClient.ConfirmERP(context.Background(), b)
	if err != nil {
		fmt.Printf("Error confirmERP: %v %v", err, res)
		return Commande{}, err
	}
	return commande, nil
}

func (commande *Commande) ToCommandePB() *commandepb.Commande {
	commandeResp := &commandepb.Commande{}
	if commande.ID != primitive.NilObjectID {
		commandeResp.ID = commande.ID.Hex()
	}
	if commande.Panier != nil {
		commandeResp.Panier = commande.Panier
	}
	if commande.Ref != "" {
		commandeResp.Ref = commande.Ref
	}
	if commande.Statut != "" {
		commandeResp.Statut = commande.Statut
	}
	if commande.Timestamp != 0 {
		commandeResp.Timestamp = commande.Timestamp
	}
	return commandeResp
}

func FromCommandePB(commandeProto *commandepb.Commande) (*Commande, error) {
	var commande = new(Commande)
	if commandeProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(commandeProto.ID)
		commande.ID = oid
	}
	if commandeProto.Panier != nil {
		commande.Panier = commandeProto.Panier
	}
	if commandeProto.Ref != "" {
		commande.Ref = commandeProto.Ref
	}
	if commandeProto.Statut != "" {
		commande.Statut = commandeProto.Statut
	}
	if commandeProto.Timestamp != primitive.NilObjectID.Timestamp().Unix() {
		commande.Timestamp = commandeProto.Timestamp
	}
	return commande, nil
}
