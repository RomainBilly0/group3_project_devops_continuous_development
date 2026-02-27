package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Title  string
	Names string
	State string
}

func main() {
	startServer()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := whoami{
		Title: "Group 3",
		Names: "Billy/Bussiere/Godfrin",
		State: "FR",
	}
	
	// Set content type to application/json
	response.Header().Set("Content-Type", "application/json")
	
	// Encode the struct to JSON and write it to the response
	err := json.NewEncoder(response).Encode(who)
	if err != nil {
		http.Error(response, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON encoding error:", err)
		return
	}
	fmt.Println("Endpoint Hit: /whoami", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "EfreiParis"

	fmt.Fprintf(response, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func startServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
