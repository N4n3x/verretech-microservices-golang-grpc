package documents

import (
	"context"
	"N4n3x/verretech-microservices/produit/produitpb"
	"fmt"
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
	Ref         string   `bson:"ref"`
	Description string   `bson:"description"`
	Prix        float32  `bson:"prix"`
	Photos      []*Photo `bson:"photos"`
	Stocks      []*Stock `bson:"stocks"`
	Tags        []string `bson:"tags"`
}

//InsertOne inserts one post in the database
func (produit *Produit) InsertOne(db mongo.Database) (primitive.ObjectID, error) {
	fmt.Printf("befor insert %+v\n",produit)
	collection := db.Collection(produitCollection)
	result, err := collection.InsertOne(context.Background(), produit)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

//FindOne returns the product with the specified ID from the database
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

//Update updates the specified produit within the database
func (produit *Produit) Update(db mongo.Database) error {
	collection := db.Collection(produitCollection)
	update := bson.M{
		"$set": bson.M{
			"ref":		produit.Ref,
			"description": produit.Description,
			"prix": produit.Prix,
			"photos": produit.Photos,
			"stocks": produit.Stocks,
			"tags": produit.Tags,
		},
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": produit.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

//FromProduitPB parses a produit defined by the protobuff into a mongo produit document
func FromProduitPB(produitProto *produitpb.Produit) (*Produit, error) {
	fmt.Printf("%+v\n",produitProto)
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
	produit.Description = produitProto.Description
	produit.Prix = produitProto.Prix
	produit.Tags = produitProto.Tags
	// produit := &Produit{
	// 	Ref: 			produitProto.Ref,
	// 	Description: 	produitProto.Description,
	// 	Prix: 			produitProto.Prix,
	// 	Photos: 		photos,
	// 	Stocks: 		stocks,
	// 	Tags: 			produitProto.Tags,
	// }
	
	fmt.Printf("End convert %+v\n",produit)
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
		l := &produitpb.Localisation{Adresse: e.PointRetrait.Localisation.Adresse, Ville: e.PointRetrait.Localisation.Ville, Cp: e.PointRetrait.Localisation.Cp}
		pr := &produitpb.PointRetrait{Nom: e.PointRetrait.Nom, Localisation: l}
		s := &produitpb.Stock{PointRetrait: pr, Qte: e.Qte}
		stocks = append(stocks, s)
	}

	return &produitpb.Produit{
		ID:       		produit.ID.Hex(),
		Ref:			produit.Ref,
		Description: 	produit.Description,
		Prix: 			produit.Prix,
		Photos: 		photos,
		Stocks: 		stocks,
		Tags: 			produit.Tags,
	}
}