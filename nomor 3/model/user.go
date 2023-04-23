package model

type User struct {
	ID          string `json:"id" bson:"_id"`
	Email       string `json:"email" bson:"email"`
	Name        string `json:"name" bson:"name"`
	Personality string `json:"personality" bson:"personality"`
	NameHash    int    `json:"namehash" bson:"namehash"`
	Password    string `json:"password" bson:"password"`
}
