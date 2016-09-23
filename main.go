package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/signup", signUpHandler)
	router.GET("/users_list", userListHandler)

	port := ":9090"
	fmt.Println("Starting server on ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
