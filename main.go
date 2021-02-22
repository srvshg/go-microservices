package main

import (
	"log"
	"microserv/handlers"
	"microserv/models"
	"microserv/repository"
	"net/http"
)

const message = "Hello World"

func main() {
	repository.AddProduct(models.Product{
		Name:      "Milk",
		UnitPrice: 5.00,
	})

	repository.AddProduct(models.Product{
		Name:      "Bread",
		UnitPrice: 4.00,
	})

	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
