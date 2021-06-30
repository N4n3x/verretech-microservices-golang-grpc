package documents

import (
	"context"
	"verretech-microservices/panier/panierpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Nom de la collection
const panierCollection = "panier"

type Article struct {
	ProduitRef string `bson:"produitRef"`
	Qte        int32  `bson:"quantite"`
}

type Panier struct {
	ID            *primitive.ObjectID `bson:"_id"`
	UtilisateurID primitive.ObjectID  `bson:"utilisateurID"`
	Articles      []*Article          `bson:"articles"`
}

//@return Panier if ID match
func (panier *Panier) FindByID(db mongo.Database, id primitive.ObjectID) {
	collection := db.Collection(panierCollection)
	query := bson.M{"utilisateurID": id}
	collection.FindOne(context.Background(), query).Decode(panier)
}

//Update updates the specified Panier within the database
func (panier *Panier) Update(db mongo.Database) (int, error) {
	opts := options.Update().SetUpsert(true)
	collection := db.Collection(panierCollection)
	update := bson.M{
		"$set": bson.M{
			"utilisateurID": panier.UtilisateurID,
			"articles":      panier.Articles,
		},
	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"utilisateurID": panier.UtilisateurID}, update, opts)
	if err != nil {
		return 0, err
	}
	return int(res.ModifiedCount), nil
}

func (panier *Panier) ToPanierPB() *panierpb.Panier {

	panierResp := &panierpb.Panier{}
	if panier.ID != nil {
		panierResp.ID = panier.ID.Hex()
	}
	panierResp.UtilisateurID = panier.UtilisateurID.Hex()
	if panier.Articles != nil {
		for _, v := range panier.Articles {
			artpb := panierpb.Article{
				Qte:        v.Qte,
				ProduitRef: v.ProduitRef,
			}
			panierResp.Article = append(panierResp.Article, &artpb)
		}
	}
	return panierResp
}

func FromPanierPB(panierProto *panierpb.Panier) (*Panier, error) {
	uid, _ := primitive.ObjectIDFromHex(panierProto.UtilisateurID)
	panier := &Panier{
		UtilisateurID: uid,
	}
	if panierProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(panierProto.ID)
		panier.ID = &oid
	}
	for _, v := range panierProto.Article {
		art := &Article{
			Qte:        v.Qte,
			ProduitRef: v.ProduitRef,
		}
		panier.Articles = append(panier.Articles, art)
	}
	return panier, nil
}
