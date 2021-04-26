package documents

import (
	"context"
	"N4n3x/verretech-microservices/produit/produitpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//produitCollection is the name of the collection storing our blog documents within the mongo database
const produitCollection = "produit"

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
	ID			primitive.ObjectID `bson:"_id,omitempty"`
	Ref         string   `bson:"ref,omitempty"`
	Description string   `bson:"description"`
	Prix        float32  `bson:"prix"`
	Photos      []*Photo `bson:"photos"`
	Stocks      []*Stock `bson:"stocks"`
	Tags        []string `bson:"tags"`
}

//InsertOne inserts one post in the database
func (produit *Produit) InsertOne(db mongo.Database) (primitive.ObjectID, error) {
	collection := db.Collection(produitCollection)
	result, err := collection.InsertOne(context.Background(), produit)
	if err != nil {
		return "Erreur", err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

//FindOne returns the post with the specified ID from the database
func (produit *Produit) FindOne(db mongo.Database) error {
	collection := db.Collection(produitCollection)
	filter := bson.M{"ref": produit.Ref}

	err := collection.FindOne(context.Background(), filter).Decode(post)
	if err != nil {
		return err
	}

	return nil
}

//Find returns a cursor pointin to all the posts in the db
func Find(db mongo.Database) (*mongo.Cursor, error) {
	collection := db.Collection(produitCollection)
	return collection.Find(context.Background(), bson.D{{}})
}

//Update updates the specified post within the database
func (produit *Produit) Update(db mongo.Database) error {
	collection := db.Collection(produitCollection)
	update := bson.M{
		"$set": bson.M{
			"ref":		produit.Ref,
			"description": produit.Description,
			"prix": produit.Prix,
			"photos": produit.Photos,
			"stocks": produit.Stock,
			"tags": produit.Tags
		},
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": produit.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

//FromPostPB parses a post defined by the protobuff into a mongo post document
func FromProduitPB(produitProto *produitpb.Produit) (*Produit, error) {
	oid, err := primitive.ObjectIDFromHex(produitProto.ID)
	if err != nil {
		return nil, err
	}

	return &Produit{
		ID:       		oid,
		Ref: 			produitProto.Ref,
		Description: 	produitProto.Description,
		Prix: 			produitProto.Prix,
		Photos: 		produitProto.Photos,
		Stocks: 		produitProto.Stocks,
		Tags: 			produitProto.Tags
	}, nil
}

//ToPostPB parses a mongo post document into a post defined by the protobuff
func (produit *Produit) ToProduitPB() *produitpb.Produit {
	return &produitpb.Produit{
		ID:       		produit.ID.Hex(),
		Ref:			produit.Ref,
		Description: 	produit.Description,
		Prix: 			produit.Prix,
		Photos: 		produit.Photos,
		Stocks: 		produit.Stocks,
		Tags: 			produit.Tags
	}
}