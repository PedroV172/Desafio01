//Funções que vai receber os inputs
package person

import (
	"sort"

	"github.com/PedroV172/challenger-algorithm-go/model"
)

func OderByAge(list []model.Person) []model.Person {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	return list
}
