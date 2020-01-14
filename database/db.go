package database

import (
	"go-bot/env"
	"log"

	"github.com/nlh1996/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mgoEnv = env.GlobalData.Server
	url    = "mongodb://" + mgoEnv.MgoAddress + ":" + utils.IntToString(mgoEnv.MgoPort)
)

// Mgo .
type Mgo struct {
	client *mongo.Client
	db     *mongo.Database
}

var mgo *Mgo

// Init .
func Init() {
	mgo = &Mgo{}
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
	mgo.db = SetDB(mgoEnv.DBName)
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

func checkErr(col string, funcName string, err error) {
	log.Println(
		"mongodb", funcName,
		"failed! db:", mgoEnv.DBName,
		", Collection:", col,
		", ErrInfo:", err,
	)
}

// InsertOne .
func InsertOne(col string, data interface{}) error {
	_, err := Col(col).InsertOne(utils.GetCtx(), data)
	if err != nil {
		checkErr(col, "InsertOne", err)
	}
	return err
}

// InsertMany .
func InsertMany(col string, data []interface{}) error {
	_, err := Col(col).InsertMany(utils.GetCtx(), data)
	if err != nil {
		checkErr(col, "InsertMany", err)
	}
	return err
}

// FindOne .
func FindOne(col string, filter interface{}, obj interface{}, opts ...*options.FindOneOptions) error {
	err := Col(col).FindOne(utils.GetCtx(), filter, opts...).Decode(obj)
	if err != nil {
		checkErr(col, "FindOne", err)
	}
	return err
}

// Find .
func Find(col string, filter interface{}, res interface{}, opts ...*options.FindOptions) error {
	cursor, err := Col(col).Find(utils.GetCtx(), filter, opts...)
	if err != nil {
		checkErr(col, "Find", err)
	}
	if err := cursor.All(utils.GetCtx(), res); err != nil {
		checkErr(col, "Find.All", err)
	}
	return err
}

// DeleteOne .
func DeleteOne(col string, filter interface{}) error {
	delRes, err := Col(col).DeleteOne(utils.GetCtx(), filter)
	if err != nil {
		checkErr(col, "DeleteOne", err)
	}
	log.Printf("DeleteOne成功删除了%d条数据。\n", delRes.DeletedCount)
	return err
}

// UpdateOne .
func UpdateOne(col string, filter interface{}, update interface{}) error {
	updateRes, err := Col(col).UpdateOne(utils.GetCtx(), filter, update)
	if err != nil {
		checkErr(col, "UpdateOne", err)
	}
	log.Println("UpdateOne成功修改一条数据", *updateRes)
	return err
}

// Count .
func Count(col string, filter interface{}) (int64, error) {
	return Col(col).CountDocuments(utils.GetCtx(), filter)
}
