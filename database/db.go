package database

import (
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
	client *mongo.Client
	db     *mongo.Database
}

var mgo *Env

// Init .
func Init() {
	mgo = &Env{}
	opts := options.Client().ApplyURI(url)
	// 需要账号认证请在SetAuth中填写
	// opts = opts.SetAuth(options.Credential{})
	var err error
	mgo.client, err = mongo.NewClient(opts)
	ctxWithTimeout := utils.GetCtx()
	err = mgo.client.Connect(ctxWithTimeout)
	if err != nil {
		log.Println(err)
	}
	mgo.db = mgo.client.Database("Blocks")
}

// SetDB .
func SetDB(str string) *mongo.Database {
	if mgo.client == nil {
		return nil
	}
	return mgo.client.Database(str)
}

// GetDB .
func GetDB() *mongo.Database {
	if mgo.db != nil {
		return mgo.db
	}
	return nil
}

// Col .
func Col(col string) *mongo.Collection {
	if mgo.db != nil {
		return mgo.db.Collection(col)
	}
	return nil
}

// InsertOne .
func InsertOne(col string, data interface{}) error {
	_, err := Col(col).InsertOne(utils.GetCtx(), data)
	return err
}

// InsertMany .
func InsertMany(col string, data []interface{}) error {
	_, err := Col(col).InsertMany(utils.GetCtx(), data)
	return err
}

// FindOne .
func FindOne(col string, filter interface{}, obj interface{}, opts ...*options.FindOneOptions) error {
	err := Col(col).FindOne(utils.GetCtx(), filter, opts...).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

// FindAll .
func FindAll(col string, filter interface{}, obj interface{}) error {

	return nil
}

// DeleteOne .
func DeleteOne(col string, filter interface{}) error {
	delRes, err := Col(col).DeleteOne(utils.GetCtx(), filter)
	if err != nil {
		return err
	}
	log.Printf("DeleteOne成功删除了%d条数据。\n", delRes.DeletedCount)
	return nil
}

// UpdateOne .
func UpdateOne(col string, filter interface{}, update interface{}) error {
	updateRes, err := Col(col).UpdateOne(utils.GetCtx(), filter, update)
	if err != nil {
		log.Println(*updateRes)
		return err
	}
	return nil
}

// Count .
func Count(col string, filter interface{}) (int64, error) {
	return Col(col).CountDocuments(utils.GetCtx(), filter)
}
