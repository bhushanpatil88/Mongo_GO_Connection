package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhushanpatil88/MONGO_GO_CONNECTION/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
	}

	oid := bson.ObjectIdHex(id)
	 
	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = bson.NewObjectId()
	uc.session.DB("mongo-golang").C("users").Insert(u)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err!= nil{
		w.WriteHeader(404)
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}