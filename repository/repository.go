package repository

import (
	"microserv/models"
)

var people []models.Person
var nextID = 1

func GetDetails() []models.Person {
	return people
}

func AddPerson(person models.Person) int {
	person.ID = nextID
	nextID++
	people = append(people, person)

	return person.ID
}
