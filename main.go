package main

import (
	"net/http"

	"github.com/bhushanpatil88/MONGO_GO_CONNECTION/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	r := mux.NewRouter()
	uc := controllers.NewUserController(getSession())
	r.HandleFunc("/user/:id",uc.GetUser).Methods("GET")
	r.HandleFunc("/user",uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/:id", uc.DeleteUser).Methods("DELETE")

	defer http.ListenAndServe(":8000",r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost:27107")

	if err != nil{
		panic(err)
	}

	return s
}