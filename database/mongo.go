package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo is the struct containing the necessary methods to connect to a mongo database
type Mongo struct {
	Database *mongo.Database
	client   *mongo.Client
}

//NewMongoConnection injects the dependencies into a new Mongo struct and returns a pointer to that struct
func NewMongoConnection() *Mongo {
	return &Mongo{}
}

//ConnectToDB connects to a mongo database with the url passed as a paremeter
func (conn *Mongo) ConnectToDB(connectionURL string, dbName string) error {
	clientOptions := options.Client().ApplyURI(connectionURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	conn.Database = client.Database(dbName)
	conn.client = client

	return nil
}

//Disconnect gracefully disconnects the client from the server
func (conn *Mongo) Disconnect() error {
	if err := conn.client.Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}