package service

import (
	"net/http"
	"github.com/gorilla/mux"
)

func GetName(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.Write([]byte("{\"name\":\"good job\"}"))
}

func GetAge(w http.ResponseWriter,r *http.Request){
	age:=mux.Vars(r)["age"]
	// w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.Write([]byte("age is "+age))
}
