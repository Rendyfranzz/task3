package main

import (
	"context"
	"crud/db"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	hashTable := NewHashTable(1000)
	conn, closeConnection, err := db.Connect("mongodb+srv://rendi:Kosongan96@cluster0.qkmm8ul.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}

	defer closeConnection(context.Background())

	log.Println("Connection ready!!!")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/create/mongo", CreateUserByMongo(conn, hashTable))
	r.Get("/create_many/mongo", CreateUserManyByMongo(conn, hashTable))

	r.Get("/get_mongo/{id}", GetOneFromMongoDB(conn))
	r.Get("/get_mongo_name/{name}", GetOneFromMongoDBHashName(conn, hashTable))

	r.Get("/get_all/mongo", GetAllFromMongoDB(conn))

	r.Get("/get_match/{name}", GetMatchedPerson(conn, hashTable))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
