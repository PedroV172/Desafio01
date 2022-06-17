package person

import (
	"sort"

	"github.com/PedroV172/challenger-algorithm-desafio/model"
)

func OderByName(list []model.Person) []model.Person {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}
