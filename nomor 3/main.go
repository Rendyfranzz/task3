package main

import (
	"context"
	"crud/db"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	client := db.MgoConn()
	coll := client.Database("task3").Collection("task3")
	for i := 0; i < 3; i++ {
		filter := bson.D{{"name", "test" + strconv.Itoa(i)}}
		var result user
		err := coll.FindOne(context.TODO(), filter).Decode(&result)
		// coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			doc := user{ID: strconv.Itoa(i), Name: "test" + strconv.Itoa(i), Email: "test" + strconv.Itoa(i), Password: "test" + strconv.Itoa(i)}
			result, err := coll.InsertOne(context.TODO(), doc)
			if err != nil {
				fmt.Println("error")
			}
			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		}
		fmt.Println("user with name test " + strconv.Itoa(i) + " already exist")
	}
}
