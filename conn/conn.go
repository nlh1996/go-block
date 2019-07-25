package conn

import (
	"fmt"
	"log"
	"user-account/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url = "mongodb://212.129.149.224:9900"
)

// Env .
type Env struct {
	database   *mongo.Database
	collection *mongo.Collection
}

var mgo *Env

func init() {
	mgo = &Env{}
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctxWithTimeout := utils.GetCtx()
	err = client.Connect(ctxWithTimeout)
	if err != nil {
		fmt.Println(err)
	}
	mgo.database = client.Database("Blocks")
	mgo.collection = mgo.database.Collection("blockchain")
}

// GetCollection .
func GetCollection() *mongo.Collection {
	if mgo.collection == nil {
		log.Fatal("no collection")
	}
	return mgo.collection
}
