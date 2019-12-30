package main

import (
	"log"
	"math/rand"
	"net/http"
)

func getRandom() string {
	rnn := rand.Intn(10)
	ps := ""
	if rnn > 7 {
		ps = "true"
	} else {
		ps = "false"
	}
	return ps
}

func NewPlan1(w http.ResponseWriter, r *http.Request) {
	ps := getRandom()
	_, _ = w.Write([]byte("{\"day\":15,\"new_plan\":true,\"emergency\":" + ps + ",\"reset\":false}"))
}

func CheckNewData(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("{\"15\":[{\"startHour\":\"22\",\"startMinute\":\"4\",\"endHour\":\"23\",\"endMinute\":\"20\"}]}"))
}
func main() {
	http.HandleFunc("/new_plan/1", NewPlan1)
	http.HandleFunc("/new_plan/2", NewPlan1)
	http.HandleFunc("/new_plan/3", NewPlan1)
	http.HandleFunc("/new_plan/4", NewPlan1)
	http.HandleFunc("/new_plan/5", NewPlan1)
	http.HandleFunc("/new_plan/6", NewPlan1)
	http.HandleFunc("/new_plan/7", NewPlan1)
	http.HandleFunc("/new_plan/8", NewPlan1)

	http.HandleFunc("/check_new_data/1", CheckNewData)
	http.HandleFunc("/check_new_data/2", CheckNewData)
	http.HandleFunc("/check_new_data/3", CheckNewData)
	http.HandleFunc("/check_new_data/4", CheckNewData)
	http.HandleFunc("/check_new_data/5", CheckNewData)
	http.HandleFunc("/check_new_data/6", CheckNewData)
	http.HandleFunc("/check_new_data/7", CheckNewData)
	http.HandleFunc("/check_new_data/8", CheckNewData)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
