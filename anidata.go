package _anidata

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataAnime(uid string, anime AnimationData, views int, likes int, comments int) (InsertedID interface{}) {
	var datastatic StatisticData
	datastatic.UniversalID = uid
	datastatic.Animation = anime
	datastatic.Views = views
	datastatic.Likes = likes
	datastatic.Comments = comments
	return InsertOneDoc("HoyoAnimation", "staticdata", datastatic)
}

func GetDataAnimeWithUID(uid string) (anidata StatisticData) {
	anilist := MongoConnect("HoyoAnimation").Collection("staticdata")
	filter := bson.M{"uid": uid}
	err := anilist.FindOne(context.TODO(), filter).Decode(&anidata)
	if err != nil {
		fmt.Printf("getDataAnimeWithUID: %v\n", err)
	}
	return anidata
}

func GetDataAnimeID(id string) (anidata StatisticData) {
	anilist := MongoConnect("HoyoAnimation").Collection("staticdata")
	filter := bson.M{"_id": id}
	err := anilist.FindOne(context.TODO(), filter).Decode(&anidata)
	if err != nil {
		fmt.Printf("GetDataAnimeByID: %v\n", err)
	}
	return anidata
}

func GetAllDataAnime() (data []StatisticData) {
	anilist := MongoConnect("HoyoAnimation").Collection("staticdata")
	filter := bson.M{}
	cursor, err := anilist.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllData: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
