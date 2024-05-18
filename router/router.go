package router

import (
	"github.com/bhushanpatil88/MONGO_GO_CONNECTION/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{

	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")


	return r
}