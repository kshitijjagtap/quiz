package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	model "github.com/kshitijjagtap/quiz_usingreact/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var userans map[string]string
var Answer_in_db []model.Answer
var userans map[string]string = make(map[string]string, 10)
var collection *mongo.Collection

func init() {
	load()
	collectioninstance()
	loadingans()
}

func loadingans() {
	Answer_in_db = Answer_puller()
}

func load() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error while loading the env file")
	}
}

func collectioninstance() {
	dbname := os.Getenv("DB_NAME")
	collName := os.Getenv("COLLECTION_NAME")
	connectionstring := os.Getenv("DB_URI")

	clientoption := options.Client().ApplyURI(connectionstring)
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		fmt.Printf("error while creating the client option")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Ping error ")
	}
	fmt.Printf("mongodb is connected now")
	collection = client.Database(dbname).Collection(collName) // add collection
	fmt.Printf("collection instance created")

}

func Answer_f(ctx *gin.Context) { // data in the format like id and answer for that question
	var data map[string]model.Answer
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"error: while binding the user": err.Error()})
	}
	answer, ok := data["data"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err != nil {
		fmt.Printf("error while binding the answer")
	}
	userans[answer.Id] = answer.Ans
	ctx.JSON(http.StatusOK, gin.H{"message": "Answer submited"})
}

func Submit(ctx *gin.Context) {
	var MarksCount int
	for _, obj := range Answer_in_db {
		if userans[obj.Id] == obj.Ans {
			MarksCount++
		}
	}

	total := strconv.Itoa(MarksCount)
	ctx.JSON(http.StatusOK, gin.H{"You got total marks:": total})
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "while binding the user"})
	}
	user.Id = primitive.NewObjectID()
	user.Marks = total
	collection.InsertOne(context.TODO(), user)

	return
}
