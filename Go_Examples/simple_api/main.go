
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World!!!!")
}

func Handlerequests(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", helloWorld).Methods("GET")
	myrouter.HandleFunc("/users", AllUsers).Methods("GET")
	myrouter.HandleFunc("/user/{name}/{email}",NewUser).Methods("POST")
	myrouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myrouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080",myrouter))

}

func main(){
	fmt.Println("Start Go From Here")
	InitialMigration()
	Handlerequests()
}
