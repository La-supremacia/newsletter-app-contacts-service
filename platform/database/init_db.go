package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {
	mongoDBUri := os.Getenv("MONGODB_URI")
	err := mgm.SetDefaultConfig(nil, "contacts", options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to mongo on URI", mongoDBUri)
	fmt.Println("Setting Indexes Rules")
	setContactIndexes()
	fmt.Println("Finished Setting Indexes Rules")

	return err
}

func setContactIndexes() {
	mod := mongo.IndexModel{Keys: bson.D{bson.E{Key: "email", Value: 1}, bson.E{Key: "organization_id", Value: 1}}, Options: options.Index().SetUnique(true)}
	name, err := mgm.CollectionByName("contacts").Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	} else {
		fmt.Println("Created index", name)
	}
}
