//Funções que vai receber os inputs
package person

import (
	"sort"

	"github.com/PedroV172/challenger-algorithm-go/model"
)

func Person(person []model.Person) []model.Person {
	sort.Slice(person, func(i, j int) bool {
		return person[i].Age < person[j].Age
	})
	return person
}
