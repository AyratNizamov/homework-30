package main

import (
	"30New/controller"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	srv := controller.Service{}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.UserFriends)
	mux.HandleFunc("/friends", srv.MakeFriends)
	mux.HandleFunc("/delete", srv.Delete)
	mux.HandleFunc("/update", srv.UpdateAge)

	http.ListenAndServe("localhost:8080", mux)
}
