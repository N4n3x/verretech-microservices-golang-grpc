package documents

import (
	"context"

	"verretech-microservices/generic/localisationpb"
	"verretech-microservices/generic/pointRetraitpb"
	"verretech-microservices/produit/produitpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Nom de la collection
const produitCollection = "produit"

///
/// Définition d'un Produit
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

type Stock struct {
	PointRetrait *PointRetrait `bson:"pointRetrait"`
	Qte          int32         `bson:"qte"`
}

type Photo struct {
	Url string `bson:"url"`
}

type Produit struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Ref         string             `bson:"ref"`
	Nom         string             `bson:"nom"`
	Description string             `bson:"description"`
	Prix        float32            `bson:"prix"`
	Photos      []*Photo           `bson:"photos"`
	Stocks      []*Stock           `bson:"stocks"`
	Tags        []string           `bson:"tags"`
}

// InsertOne Insert un produit en base de données
// Retourne ObjectID du produit si l'insertion se passe bien, ou une erreur
func (produit *Produit) InsertOne(db mongo.Database) (primitive.ObjectID, error) {
	// fmt.Printf("befor insert %+v\n", produit)
	collection := db.Collection(produitCollection)
	result, err := collection.InsertOne(context.Background(), produit)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// FindOne  un Produit à partir de ça propriété Ref, ou
func (produit *Produit) FindOne(db mongo.Database) error {
	collection := db.Collection(produitCollection)
	filter := bson.M{"ref": produit.Ref}

	err := collection.FindOne(context.Background(), filter).Decode(produit)
	if err != nil {
		return err
	}

	return nil
}

//Find returns a cursor pointin to all the produits in the db
func Find(db mongo.Database) (*mongo.Cursor, error) {
	collection := db.Collection(produitCollection)
	return collection.Find(context.Background(), bson.D{{}})
}

//Find returns a cursor pointin to all the produits in the db
func FindByTags(db mongo.Database, tags primitive.M) (*mongo.Cursor, error) {
	collection := db.Collection(produitCollection)
	return collection.Find(context.Background(), tags)
}

//Update updates the specified produit within the database
func (produit *Produit) Update(db mongo.Database) (int, error) {
	opts := options.Update().SetUpsert(true)
	collection := db.Collection(produitCollection)
	update := bson.M{
		"$set": bson.M{
			"ref":         produit.Ref,
			"nom":         produit.Nom,
			"description": produit.Description,
			"prix":        produit.Prix,
			"photos":      produit.Photos,
			"stocks":      produit.Stocks,
			"tags":        produit.Tags,
		},
	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": produit.ID}, update, opts)
	if err != nil {
		return 0, err
	}
	return int(res.ModifiedCount), nil
}

//Update updates the specified produit within the database
func UpdateAll(db mongo.Database, produits []*Produit) (int, error) {
	collection := db.Collection(produitCollection)
	collection.DeleteMany(context.Background(), bson.M{})
	p := []interface{}{}
	for _, produit := range produits {
		p = append(p, bson.M{
			"ref":         produit.Ref,
			"nom":         produit.Nom,
			"description": produit.Description,
			"prix":        produit.Prix,
			"photos":      produit.Photos,
			"stocks":      produit.Stocks,
			"tags":        produit.Tags,
		})
	}
	res, err := collection.InsertMany(context.Background(), p)
	if err != nil {
		return 0, err
	}
	return len(res.InsertedIDs), nil
}

func (produit *Produit) Delete(db mongo.Database) (int, error) {
	collection := db.Collection(produitCollection)
	res, err := collection.DeleteOne(context.Background(), bson.M{"ref": produit.Ref})
	if err != nil {
		return 0, err
	}
	return int(res.DeletedCount), nil
}

//FromProduitPB parses a produit defined by the protobuff into a mongo produit document
func FromProduitPB(produitProto *produitpb.Produit) (*Produit, error) {
	// fmt.Printf("%+v\n", produitProto)
	var produit = new(Produit)
	var photos []*Photo
	var stocks []*Stock

	if produitProto.ID != "" {
		oid, _ := primitive.ObjectIDFromHex(produitProto.ID)
		produit.ID = oid
	}

	for _, e := range produitProto.Photos {
		p := &Photo{Url: e.Url}
		photos = append(photos, p)
	}
	produit.Photos = photos
	// fmt.Printf("%+v\n",photos)

	for _, e := range produitProto.Stocks {
		l := &Localisation{Adresse: e.PointRetrait.Localisation.Adresse, Ville: e.PointRetrait.Localisation.Ville, Cp: e.PointRetrait.Localisation.Cp}
		pr := &PointRetrait{Nom: e.PointRetrait.Nom, Localisation: l}
		s := &Stock{PointRetrait: pr, Qte: e.Qte}
		stocks = append(stocks, s)
	}
	produit.Stocks = stocks

	produit.Ref = produitProto.Ref
	produit.Nom = produitProto.Nom
	produit.Description = produitProto.Description
	produit.Prix = produitProto.Prix
	produit.Tags = produitProto.Tags

	// fmt.Printf("End convert %+v\n", produit)
	return produit, nil
}

//ToProduitPB parses a mongo produit document into a produit defined by the protobuff
func (produit *Produit) ToProduitPB() *produitpb.Produit {
	var photos []*produitpb.Photo
	for _, e := range produit.Photos {
		p := &produitpb.Photo{Url: e.Url}
		photos = append(photos, p)
	}

	var stocks []*produitpb.Stock
	for _, e := range produit.Stocks {
		l := &localisationpb.Localisation{Adresse: e.PointRetrait.Localisation.Adresse, Ville: e.PointRetrait.Localisation.Ville, Cp: e.PointRetrait.Localisation.Cp}
		pr := &pointRetraitpb.PointRetrait{Nom: e.PointRetrait.Nom, Localisation: l}
		s := &produitpb.Stock{PointRetrait: pr, Qte: e.Qte}
		stocks = append(stocks, s)
	}

	return &produitpb.Produit{
		ID:          produit.ID.Hex(),
		Ref:         produit.Ref,
		Nom:         produit.Nom,
		Description: produit.Description,
		Prix:        produit.Prix,
		Photos:      photos,
		Stocks:      stocks,
		Tags:        produit.Tags,
	}
}
