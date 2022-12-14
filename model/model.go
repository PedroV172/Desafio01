//Estrutura de dados(struct)
package model

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      int    `json:"age"`
}
