package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/PedroV172/challenger-algorithm-desafio/model"
	"github.com/PedroV172/challenger-algorithm-desafio/person"
)

func main() {
	list := []model.Person{}
	fileList, err := os.Open("list_person.json")
	if err != nil {
		log.Fatal(err)
	}

	fileListByte, err := ioutil.ReadAll(fileList)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileListByte, &list)
	if err != nil {
		log.Fatal(err)
	}

	for index, person := range person.OderByName(list) {
		result := fmt.Sprintf("%v) %v, idade %v", index+1, person.Name, person.Age)
		fmt.Println(result)
	}
}