
package main

import (
	"SMini/HandleRequests"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleURLs(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/new_plan/{ID}", HandleRequests.HandlePlanStatus).Methods("GET")
	myrouter.HandleFunc("/check_new_data/{ID}", HandleRequests.HandleTimeStatus).Methods("GET")
	myrouter.HandleFunc("/check_new_data/{ID}/{day}", HandleRequests.DeleteTimeStatus).Methods("DELETE")
	myrouter.HandleFunc("/check_new_data/{ID}/{day}/{StartHour}/{StartMinute}/{EndHour}/{EndMinute}", HandleRequests.UpdateTimeStatus).Methods("PUT")
	myrouter.HandleFunc("/check_new_data/{ID}/{day}/{StartHour}/{StartMinute}/{EndHour}/{EndMinute}", HandleRequests.CreateTimeStatus).Methods("POST")
	myrouter.HandleFunc("/new_plan/{ID}/{emergency}/{reset}", HandleRequests.UpdatePlanStatus).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080",myrouter))

}

func main(){
	HandleRequests.SqlMigration()
	go HandleRequests.GetDay()
	HandleURLs()
}
