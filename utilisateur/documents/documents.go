package documents

import (
	"context"

	"verretech-microservices/generic/localisationpb"
	"verretech-microservices/generic/pointRetraitpb"
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
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Nom            string             `bson:"nom"`
	Prenom         string             `bson:"prenom"`
	Mail           string             `bson:"mail"`
	HashMotDePasse string             `bson:"hashmotdepasse"`
	Preferences    *Preferences       `bson:"preferences"`
	Permission     []string           `bson:"permission"`
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

// FindOne  un Utilisateur à partir de ça propriété Mail, ou
func (utilisateur *Utilisateur) FindOne(db mongo.Database) error {
	collection := db.Collection(utilisateurCollection)
	filter := bson.M{"mail": utilisateur.Mail}

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

	utilisateur := &Utilisateur{
		Nom:            utilisateurProto.Nom,
		Prenom:         utilisateurProto.Prenom,
		Mail:           utilisateurProto.Mail,
		HashMotDePasse: utilisateurProto.HashMotDePasse,
		Preferences: &Preferences{
			Localisation: &Localisation{
				Adresse: utilisateurProto.Preferences.Localisation.Adresse,
				Ville:   utilisateurProto.Preferences.Localisation.Ville,
				Cp:      utilisateurProto.Preferences.Localisation.Cp,
			},
			PointRetrait: &PointRetrait{
				Nom: utilisateurProto.Preferences.PointRetrait.Nom,
				Localisation: &Localisation{
					Adresse: utilisateurProto.Preferences.PointRetrait.Localisation.Adresse,
					Ville:   utilisateurProto.Preferences.PointRetrait.Localisation.Ville,
					Cp:      utilisateurProto.Preferences.PointRetrait.Localisation.Cp,
				},
			},
		},
		Permission: utilisateurProto.Permission,
	}

	if utilisateurProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(utilisateurProto.ID)
		utilisateur.ID = oid
	}

	// fmt.Printf("End convert %+v\n", utilisateur)
	return utilisateur, nil
}

//ToProduitPB parses a mongo produit document into a produit defined by the protobuff
func (utilisateur *Utilisateur) ToUtilisateurPB() *utilisateurpb.Utilisateur {
	return &utilisateurpb.Utilisateur{
		ID:             utilisateur.ID.Hex(),
		Nom:            utilisateur.Nom,
		Prenom:         utilisateur.Prenom,
		Mail:           utilisateur.Mail,
		HashMotDePasse: utilisateur.HashMotDePasse,
		Preferences: &utilisateurpb.Preferences{
			Localisation: &localisationpb.Localisation{
				Adresse: utilisateur.Preferences.Localisation.Adresse,
				Ville:   utilisateur.Preferences.Localisation.Ville,
				Cp:      utilisateur.Preferences.Localisation.Cp,
			},
			PointRetrait: &pointRetraitpb.PointRetrait{
				Nom: utilisateur.Preferences.PointRetrait.Nom,
				Localisation: &localisationpb.Localisation{
					Adresse: utilisateur.Preferences.PointRetrait.Localisation.Adresse,
					Ville:   utilisateur.Preferences.PointRetrait.Localisation.Ville,
					Cp:      utilisateur.Preferences.PointRetrait.Localisation.Cp,
				},
			},
		},
	}
}
