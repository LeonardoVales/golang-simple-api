package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api está rodando")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

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
