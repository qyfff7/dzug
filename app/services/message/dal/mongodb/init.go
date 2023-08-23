package mongodb

import (
	"context"
	"dzug/conf"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func Init() error {
	var err error
	uri := fmt.Sprintf("mongodb://%s", conf.Config.MongoDbConfig.Addr)
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Print(err)
		return err
	}
	//2.选择数据库 my_db
	db = client.Database("message")

	//3.选择表 my_collection
	collection = db.Collection("message")

	return nil

	//index := mongo.IndexModel{
	//	Keys:    bson.M{"created_time": 1},
	//	Options: options.Index().SetExpireAfterSeconds(int32(time.Now().Add(time.Hour * 24).Unix())),
	//}
	//_, err = collection.Indexes().CreateOne(context.Background(), index)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
