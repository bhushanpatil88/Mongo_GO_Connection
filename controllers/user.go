package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bhushanpatil88/MONGO_GO_CONNECTION/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)




var collection *mongo.Collection

func init(){
	godotenv.Load()
	connectionString := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("COLLECTION_NAME")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successfull")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reference is ready")
}

func insertOne(user models.User){
	inserted, err := collection.InsertOne(context.Background(), user)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted data with id: ", inserted.InsertedID)
}

func updateOne(id string){
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id":oid}
	update := bson.M{"$set":bson.M{"age":100}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}
}

func deleteOne(id string){
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id":oid}
	
	_, err := collection.DeleteOne(context.Background(), filter)
	
	if err != nil{
		log.Fatal(err)
	}
}

func getAllUsers() [] models.User{
	cur, err := collection.Find(context.Background(), primitive.M{})

	if err != nil{
		log.Fatal(err)
	}
	var users []models.User

	for  cur.Next(context.Background()){
		var user models.User

		err := cur.Decode(&user)
		if err != nil{
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())

	return users
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	allUsers := getAllUsers()

	json.NewEncoder(w).Encode(allUsers)

}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user models.User 
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertOne(user)

	json.NewEncoder(w).Encode(user)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updateOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}