package main

import (
	"crud/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Map map[string]interface{}

func CreateUserByMongo(conn *mongo.Client, t *HashTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}

		defer r.Body.Close()

		user.ID = xid.New().String()
		user.NameHash = t.Hash(user.Name)
		user.Personality = Personality()

		collection := conn.Database("task3").Collection("task3")
		result, err := collection.InsertOne(r.Context(), &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}

func CreateUserManyByMongo(conn *mongo.Client, t *HashTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query().Get("count")
		count, _ := strconv.Atoi(qp)

		users := make([]any, count)

		for i := 0; i < count; i++ {
			user := model.User{
				ID:          xid.New().String(),
				Name:        "name_" + strconv.Itoa(i),
				Email:       "email_" + strconv.Itoa(i),
				NameHash:    t.Hash("name_" + strconv.Itoa(i)),
				Personality: Personality(),
				Password:    "password_" + strconv.Itoa(i),
			}

			users[i] = user
		}

		collection := conn.Database("task3").Collection("task3")
		results, err := collection.InsertMany(r.Context(), users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Map{"error inserting": err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(results)
	}
}

func GetAllFromMongoDB(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := conn.Database("task3").Collection("task3")
		rows, err := collection.Find(r.Context(), bson.M{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Map{"error finding": err.Error()})
			return
		}

		var users []model.User

		for rows.Next(r.Context()) {
			var user model.User

			err := rows.Decode(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Map{"error streaming": err.Error()})
				return
			}

			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	}
}

func GetOneFromMongoDB(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyToFind := chi.URLParam(r, "id")

		collection := conn.Database("task3").Collection("task3")

		var user model.User

		err := collection.FindOne(r.Context(), bson.M{"_id": keyToFind}).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func GetOneFromMongoDBHashName(conn *mongo.Client, t *HashTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyToFind := chi.URLParam(r, "name")

		collection := conn.Database("task3").Collection("task3")

		var user model.User
		namehash := t.Hash(keyToFind)
		err := collection.FindOne(r.Context(), bson.M{"namehash": namehash}).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func GetMatchedPerson(conn *mongo.Client, t *HashTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyToFind := chi.URLParam(r, "name")
		collection := conn.Database("task3").Collection("task3")
		var user model.User
		namehash := t.Hash(keyToFind)
		err := collection.FindOne(r.Context(), bson.M{"namehash": namehash}).Decode(&user)
		fmt.Println(user.Personality)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}
		var personality string = MatchPerson(user.Personality)
		var users []model.User
		rows, err := collection.Find(r.Context(), bson.D{{"personality", personality}})
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}
		for rows.Next(r.Context()) {
			var user model.User

			err := rows.Decode(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Map{"error streaming": err.Error()})
				return
			}

			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	}
}
