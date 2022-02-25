package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")
}

func NewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Users Endpoint Hit")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Users Endpoint Hit")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Users Endpoint Hit")
}
