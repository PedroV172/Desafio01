//Funções que vai receber os inputs
package person

import "fmt"

type People struct {
	Name string
	Age int
}

func Person(){
    var u People
	u.Name = "Pedro"
	u.Age = 19
	fmt.Println(u)
}