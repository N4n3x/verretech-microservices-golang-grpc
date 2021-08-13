package documents

import (
	"context"
	"errors"

	"verretech-microservices/utilisateur/utilisateurpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Nom de la collection
const utilisateurCollection = "utilisateur"

///
/// Définition d'un Utilisateur
///
type Localisation struct {
	Adresse string `bson:"adresse"`
	Ville   string `bson:"ville"`
	Cp      string `bson:"cp"`
}

type PointRetrait struct {
	Nom          string        `bson:"nom"`
	Localisation *Localisation `bson:"localisation"`
}

type Preferences struct {
	Localisation *Localisation `bson:"localisation"`
	PointRetrait *PointRetrait `bson:"pointretrait"`
}

type Utilisateur struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty"`
	Nom            *string             `bson:"nom"`
	Prenom         *string             `bson:"prenom"`
	Mail           *string             `bson:"mail"`
	HashMotDePasse *string             `bson:"hashmotdepasse"`
	Preferences    *Preferences        `bson:"preferences"`
	Permission     *[]string           `bson:"permission"`
}

// InsertOne Ajoute un utilisateur en base de données
// Retourne ObjectID de l'utilisateur si l'insertion se passe bien, ou une erreur
func (utilisateur *Utilisateur) InsertOne(db mongo.Database) (primitive.ObjectID, error) {
	// fmt.Printf("befor insert %+v\n", produit)
	collection := db.Collection(utilisateurCollection)
	result, err := collection.InsertOne(context.Background(), utilisateur)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// FindOne  retourne un Utilisateur à partir de ça propriété Mail, ou ID
func (utilisateur *Utilisateur) FindOne(db mongo.Database) error {
	collection := db.Collection(utilisateurCollection)
	var filter bson.M
	if utilisateur.ID != nil && utilisateur.ID != &primitive.NilObjectID {
		filter = bson.M{"_id": utilisateur.ID}

	} else if utilisateur.Mail != nil {
		filter = bson.M{"mail": utilisateur.Mail}

	} else {
		return errors.New("Utilisateur erreur, pas de filtre")
	}

	err := collection.FindOne(context.Background(), filter).Decode(utilisateur)
	if err != nil {
		return err
	}

	return nil
}

//Find retourne un cursor vers tous les utilisateur
func Find(db mongo.Database) (*mongo.Cursor, error) {
	collection := db.Collection(utilisateurCollection)
	return collection.Find(context.Background(), bson.D{{}})
}

//Update met à jour un Utilisateur par son ID ou ajoute en base
func (utilisateur *Utilisateur) Update(db mongo.Database) (int, error) {
	opts := options.Update().SetUpsert(true)
	collection := db.Collection(utilisateurCollection)
	update := bson.M{
		"$set": bson.M{
			"nom":            utilisateur.Nom,
			"prenom":         utilisateur.Prenom,
			"mail":           utilisateur.Mail,
			"hashmotdepasse": utilisateur.HashMotDePasse,
			"preferences":    utilisateur.Preferences,
			"permission":     utilisateur.Permission,
		},
	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": utilisateur.ID}, update, opts)
	if err != nil {
		return 0, err
	}
	return int(res.ModifiedCount), nil
}

///Delete Supprime un Utilisateur par son mail
func (utilisateur *Utilisateur) Delete(db mongo.Database) (int, error) {
	collection := db.Collection(utilisateurCollection)
	res, err := collection.DeleteOne(context.Background(), bson.M{"mail": utilisateur.Mail})
	if err != nil {
		return 0, err
	}
	return int(res.DeletedCount), nil
}

//FromUtilisateurPB Instancie un Utilisateur à partir d'un Utilisateur ProtoBuf
func FromUtilisateurPB(utilisateurProto *utilisateurpb.Utilisateur) (*Utilisateur, error) {
	// fmt.Printf("%+v\n", utilisateurProto)
	///TODO: Null safety
	utilisateur := &Utilisateur{}
	if utilisateurProto.Nom != "" {
		utilisateur.Nom = &utilisateurProto.Nom
	}
	if utilisateurProto.Prenom != "" {
		utilisateur.Prenom = &utilisateurProto.Prenom
	}
	if utilisateurProto.Mail != "" {
		utilisateur.Mail = &utilisateurProto.Mail
	}
	if utilisateurProto.HashMotDePasse != "" {
		utilisateur.HashMotDePasse = &utilisateurProto.HashMotDePasse
	}
	if utilisateurProto.Permission != nil {
		utilisateur.Permission = &utilisateurProto.Permission
	}
	if utilisateurProto.Preferences != nil {
		if utilisateurProto.Preferences.Localisation != nil {
			utilisateur.Preferences.Localisation.Adresse = utilisateurProto.Preferences.Localisation.Adresse
			utilisateur.Preferences.Localisation.Ville = utilisateurProto.Preferences.Localisation.Ville
			utilisateur.Preferences.Localisation.Cp = utilisateurProto.Preferences.Localisation.Cp
		}
		if utilisateurProto.Preferences.PointRetrait != nil {
			utilisateur.Preferences.PointRetrait.Nom = utilisateurProto.Preferences.PointRetrait.Nom
			if utilisateurProto.Preferences.PointRetrait.Localisation != nil {
				utilisateur.Preferences.PointRetrait.Localisation.Adresse = utilisateurProto.Preferences.PointRetrait.Localisation.Adresse
				utilisateur.Preferences.PointRetrait.Localisation.Ville = utilisateurProto.Preferences.PointRetrait.Localisation.Ville
				utilisateur.Preferences.PointRetrait.Localisation.Cp = utilisateurProto.Preferences.PointRetrait.Localisation.Cp
			}
		}
	}

	if utilisateurProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(utilisateurProto.ID)
		utilisateur.ID = &oid
	}

	// fmt.Printf("End convert %+v\n", utilisateur)
	return utilisateur, nil
}

//ToProduitPB parses a mongo produit document into a produit defined by the protobuff
func (utilisateur *Utilisateur) ToUtilisateurPB() *utilisateurpb.Utilisateur {
	utilisateurResp := &utilisateurpb.Utilisateur{}
	if utilisateur.ID != nil {
		utilisateurResp.ID = utilisateur.ID.Hex()
	}
	if utilisateur.Nom != nil {
		utilisateurResp.Nom = *utilisateur.Nom
	}
	if utilisateur.Prenom != nil {
		utilisateurResp.Prenom = *utilisateur.Prenom
	}
	if utilisateur.Mail != nil {
		utilisateurResp.Mail = *utilisateur.Mail
	}
	if utilisateur.HashMotDePasse != nil {
		utilisateurResp.HashMotDePasse = *utilisateur.HashMotDePasse
	}

	if utilisateur.Preferences != nil {
		if utilisateur.Preferences.Localisation != nil {
			utilisateurResp.Preferences.Localisation.Adresse = utilisateur.Preferences.Localisation.Adresse
			utilisateurResp.Preferences.Localisation.Ville = utilisateur.Preferences.Localisation.Ville
			utilisateurResp.Preferences.Localisation.Cp = utilisateur.Preferences.Localisation.Cp
		}
		if utilisateur.Preferences.PointRetrait != nil {
			utilisateurResp.Preferences.PointRetrait.Nom = utilisateur.Preferences.PointRetrait.Nom
			if utilisateur.Preferences.PointRetrait.Localisation != nil {
				utilisateurResp.Preferences.PointRetrait.Localisation.Adresse = utilisateur.Preferences.PointRetrait.Localisation.Adresse
				utilisateurResp.Preferences.PointRetrait.Localisation.Ville = utilisateur.Preferences.PointRetrait.Localisation.Ville
				utilisateurResp.Preferences.PointRetrait.Localisation.Cp = utilisateur.Preferences.PointRetrait.Localisation.Cp
			}
		}
	}

	return utilisateurResp
}
