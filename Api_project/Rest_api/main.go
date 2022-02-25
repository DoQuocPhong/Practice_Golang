package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/fatih/color"
	//"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	article := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articals Endpoint")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){

	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", homePage)
	myRoute.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRoute))

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", allArticles)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}

