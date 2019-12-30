package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Articale struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Articale
func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Articale{Title:"Test Title", Desc:"Test Desc", Content:"Test Content"},
		Articale{Title:"Test Title", Desc:"Test Desc", Content:"Test Content"},
	}
	fmt.Fprintln(w,"Endpoint Hit: All Articles")
	json.NewEncoder(w).Encode(articles)
}

func TestPostArticles(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "POST TEST completed")
}

func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page Endpoint HI")
}
func HandleHomeRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", TestPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main(){
	HandleHomeRequests()
}