package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-study/web/model"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s", name)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request:", err)
	}
	user.GenerateCreateTime()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func NewHttpRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/foo", FooHandler)
	router.HandleFunc("/user", CreateUserHandler).Methods("POST")

	return router
}
