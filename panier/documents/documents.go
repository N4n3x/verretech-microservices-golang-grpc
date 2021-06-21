package documents

import (
	"context"
	"verretech-microservices/panier/panierpb"
	produitDoc "verretech-microservices/produit/documents"
	utilisateurDoc "verretech-microservices/utilisateur/documents"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Nom de la collection
const panierCollection = "panier"

type Article struct {
	Produit *produitDoc.Produit `bson:"produit"`
	Qte     int32               `bson:"quantite"`
}

type Panier struct {
	ID          primitive.ObjectID          `bson:"_id,omitempty"`
	Utilisateur *utilisateurDoc.Utilisateur `bson:"utilisateur"`
	Articles    []*Article                  `bson:"articles"`
}

//@return Panier if ID match
func FindByID(db mongo.Database, id primitive.M) *mongo.SingleResult {
	collection := db.Collection(panierCollection)
	return collection.FindOne(context.Background(), id)
}

//Update updates the specified Panier within the database
func (panier *Panier) Update(db mongo.Database) (int, error) {
	opts := options.Update().SetUpsert(true)
	collection := db.Collection(panierCollection)
	update := bson.M{
		"$set": bson.M{
			"utilisateur": panier.Utilisateur,
			"articles":    panier.Articles,
		},
	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": panier.ID}, update, opts)
	if err != nil {
		return 0, err
	}
	return int(res.ModifiedCount), nil
}

func (panier *Panier) ToPanierPB() *panierpb.Panier {
	panierResp := &panierpb.Panier{
		ID: panier.ID.Hex(),
		// Utilisateur:    panier.Utilisateur,
		// Articles:       panier.Articles,
	}

	if panier.Utilisateur != nil {
		panierResp.Utilisateur = panier.Utilisateur.ToUtilisateurPB()
	}
	if panier.Articles != nil {
		for _, v := range panier.Articles {
			artpb := panierpb.Article{
				Qte:     v.Qte,
				Produit: v.Produit.ToProduitPB(),
			}
			panierResp.Article = append(panierResp.Article, &artpb)
		}
	}
	return panierResp
}

func FromPanierPB(panierProto *panierpb.Panier) (*Panier, error) {
	utilisateur, err := utilisateurDoc.FromUtilisateurPB(panierProto.Utilisateur)
	if err != nil {
		return nil, err
	}
	panier := &Panier{
		Utilisateur: utilisateur,
	}
	for _, v := range panierProto.Article {
		p, e := produitDoc.FromProduitPB(v.Produit)
		if e != nil {
			return nil, e
		}
		art := &Article{
			Qte:     v.Qte,
			Produit: p,
		}
		panier.Articles = append(panier.Articles, art)
	}
	return panier, nil
}
