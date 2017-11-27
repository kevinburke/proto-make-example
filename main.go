package main

import (
	"net/http"

	proto "github.com/golang/protobuf/proto"
	"github.com/kevinburke/handlers"
	"github.com/kevinburke/proto-make-example/assets"
	"github.com/kevinburke/rest"
)

// This is an example application that serves Protobuf data and also data
// compiled from templates.

func main() {
	// Write a protobuf User to the client.
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		u := &User{Id: 1, Email: "test@example.com", Name: "test user"}
		data, err := proto.Marshal(u)
		if err != nil {
			rest.ServerError(w, r, err)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(data)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := assets.Asset("static/index.html")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})
	http.ListenAndServe(":8080", handlers.Log(http.DefaultServeMux))
}
