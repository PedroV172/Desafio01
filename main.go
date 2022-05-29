package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PedroV172/challenger-algorithm-go/model"
	"github.com/PedroV172/challenger-algorithm-go/person"
)

func main() {
	list := []model.Person{}
	fileList, err := os.Open("list_person.json")
	if err != nil {
		panic(errors.New("error open file json" + err.Error()))
	}

	fileListByte, err := ioutil.ReadAll(fileList)
	if err != nil {
		panic(errors.New("error convert file in bytes" + err.Error()))
	}

	err = json.Unmarshal(fileListByte, &list)
	if err != nil {
		panic(errors.New("error unmarshal" + err.Error()))
	}

	for index, person := range person.OderByAge(list) {
		result := fmt.Sprintf("%v) %v, idade %v", index+1, person.Name, person.Age)
		fmt.Println(result)
	}
}
