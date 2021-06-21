package documents

import (
	"verretech-microservices/panier/panierpb"
	produitDoc "verretech-microservices/produit/documents"
	utilisateurDoc "verretech-microservices/utilisateur/documents"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (panier *Panier) ToPanierPB() *panierpb.Panier {
	panierResp := &panierpb.Panier{
		ID:             panier.ID.Hex(),
		Nom:            panier.Nom,
		Prenom:         panier.Prenom,
		Mail:           panier.Mail,
		HashMotDePasse: panier.HashMotDePasse,
	}

	if panier.Preferences != nil {
		if panier.Preferences.Localisation != nil {
			panierResp.Preferences.Localisation.Adresse = panier.Preferences.Localisation.Adresse
			panierResp.Preferences.Localisation.Ville = panier.Preferences.Localisation.Ville
			panierResp.Preferences.Localisation.Cp = panier.Preferences.Localisation.Cp
		}
		if panier.Preferences.PointRetrait != nil {
			panierResp.Preferences.PointRetrait.Nom = panier.Preferences.PointRetrait.Nom
			if panier.Preferences.PointRetrait.Localisation != nil {
				panierResp.Preferences.PointRetrait.Localisation.Adresse = panier.Preferences.PointRetrait.Localisation.Adresse
				panierResp.Preferences.PointRetrait.Localisation.Ville = panier.Preferences.PointRetrait.Localisation.Ville
				panierResp.Preferences.PointRetrait.Localisation.Cp = panier.Preferences.PointRetrait.Localisation.Cp
			}
		}
	}

	return panierResp
}
