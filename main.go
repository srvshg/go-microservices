package main

import (
	"log"
	"microserv/handlers"
	"microserv/models"
	"microserv/repository"
	"net/http"
)

func main() {
	repository.AddPerson(models.Person{
		Name: "Mike",
		Age:  27,
	})

	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
