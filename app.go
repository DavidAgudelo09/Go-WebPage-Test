package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Users struct {
	Name   string
	Skills string
	Age    int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := Users{"David Agudelo", "Software Engineer", 24}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func main() {

	http.HandleFunc("/", Index)

	// Crea el servidor
	fmt.Println("The server is running on port 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("https://davidagudelo09.github.io/Go-WebPage-Test/", nil)
}
