package main

import (
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	article := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello Worlddddddeeee"},
	}

	fmt.Println("Endpoint Hit: All Articals Endpoint")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}

