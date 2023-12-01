package connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Database {

	option := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	var db *mongo.Database = client.Database("dpw")

	return db

}
