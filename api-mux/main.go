package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("tá rodando")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id:   1,
			Name: "Leonardo Vales",
		},
		{
			Id:   2,
			Name: "Pollyana",
		},
		{
			Id:   3,
			Name: "Joyce Amanda",
		},
	})
}
