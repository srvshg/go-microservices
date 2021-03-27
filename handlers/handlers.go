package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"microserv/models"
	"microserv/repository"
	"net/http"
)

// handler is here
func HandleRequest(w http.ResponseWriter, r *http.Request) {

	log.Println("Request recived: ", r.Method)

	switch r.Method {
	case http.MethodGet:
		display(w, r)
		break
	case http.MethodPost:
		add(w, r)
		break
	default:
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}

}

func display(w http.ResponseWriter, r *http.Request) {
	products := repository.GetDetails()
	json, _ := json.Marshal(products) //ignored the error

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)

	log.Println("Response returned: ", 200)

}

func add(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var person models.Person
	err := json.Unmarshal(payload, &person)
	if err != nil || person.Name == "" || person.Age == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	person.ID = repository.AddPerson(person)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(201)

	json, _ := json.Marshal(person)
	w.Write(json)

	log.Println("Response returned: ", 201)
}
