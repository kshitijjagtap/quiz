package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	Id  string `json:"id"`
	Ans string `json:"ans"`
}

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username"`
	Marks    string             `json:"marks"`
	Time     time.Time          `json:"time"`
}
