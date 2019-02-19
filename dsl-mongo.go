package main

import (
	"context"
	_ "fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	//	"reflect"
	"time"
)

func init() {
	client, _ := mongo.NewClient("mongodb://heroku_06k6mm29:o3kt6mvjjg92g6gj8v9mqdjp5h@ds037387.mlab.com:37387/heroku_06k6mm29")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Hour)
	client.Connect(ctx)

	DslFunctions["mongoGet"] = func(container map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].rawArg.(string)
		collection := client.Database("heroku_06k6mm29").Collection(collectionName)
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		records := []map[string]interface{}{}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result map[string]interface{}
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, result)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		return records, nil
	}

	DslFunctions["mongoInsert"] = func(container map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].rawArg.(string)
		obj, err := args[1].Evaluate(container)
		if err != nil {
			return nil, err
		}
		collection := client.Database("heroku_06k6mm29").Collection(collectionName)
		res, err := collection.InsertOne(ctx, obj)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	DslFunctions["mongoReplace"] = func(container map[string]interface{}, args ...Argument) (interface{}, error) {
		collectionName := args[0].rawArg.(string)
		obj, err := args[1].Evaluate(container)
		if err != nil {
			return nil, err
		}
		collection := client.Database("heroku_06k6mm29").Collection(collectionName)
		res := collection.FindOneAndReplace(ctx, map[string]interface{}{"_id": (obj.(map[string]interface{}))["_id"]}, obj)
		return res, nil
	}
}