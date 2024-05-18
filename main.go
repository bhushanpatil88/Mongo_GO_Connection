package main

import (
	"fmt"
	"net/http"

	"github.com/bhushanpatil88/MONGO_GO_CONNECTION/router"
)

func main() {
	r := router.Router()
	fmt.Println("Mongo API")
	fmt.Println("Server is getting started")
	defer http.ListenAndServe(":9000",r)
}

