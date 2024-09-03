package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Database struct {
	client *mongo.Client
}

func Connect(uri string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	log.Println("Conectado a MongoDB")
	return &Database{client: client}, nil
}

func (db *Database) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error desconectando de MongoDB: %v", err)
	}
}
