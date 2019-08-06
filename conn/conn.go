package conn

import (
	"fmt"
	"log"
	"user-account/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url = "mongodb://localhost:27017"
)

// Env .
type Env struct {
	database   *mongo.Database
	collection *mongo.Collection
}

var mgo *Env

// Init .
func Init() {
	mgo = &Env{}
	opts := options.Client().ApplyURI(url)
	// 需要账号认证请在SetAuth中填写
	// opts = opts.SetAuth(options.Credential{})
	client, err := mongo.NewClient(opts)
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

// SetCollection .
func SetCollection(str string) {
	if mgo != nil {
		mgo.collection = mgo.database.Collection(str)
	}
}