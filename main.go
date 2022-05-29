package main

import (
	"fmt"

	"github.com/PedroV172/challenger-algorithm-go/model"
	"github.com/PedroV172/challenger-algorithm-go/person"
)

func main() {
	persoa := []model.Person{}

	persoa = append(persoa, model.Person{
		Name: "witalo",
		Age:  30,
	})

	persoa = append(persoa, model.Person{
		Name: "Pedro",
		Age:  18,
	})

	fmt.Println(person.Person(persoa))
}
