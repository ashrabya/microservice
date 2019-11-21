package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Account will hold
type Account struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber int64  `json:"phonenumber"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
}

var acc []Account

func main() {

	initHandler()

}

func initHandler() {

	r := mux.NewRouter()

	r.HandleFunc("/account", createAccount).Methods(http.MethodPost)
	r.HandleFunc("/account", getAccount).Methods(http.MethodGet)
	r.HandleFunc("/account", updateAccount).Methods(http.MethodPut)
	r.HandleFunc("/account", deleteAccount).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var a Account

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	acc = append(acc, a)
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "Account Created Successfully")
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func updateAccount(w http.ResponseWriter, r *http.Request) {}

func deleteAccount(w http.ResponseWriter, r *http.Request) {}
