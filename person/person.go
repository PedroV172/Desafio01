//Funções que vai receber os inputs
package person

import "fmt"

type People struct {
	Name string
	Age int
}

func Person(){
    var u1 People
	u1.Name = "João"
	u1.Age = 19
	fmt.Println(u1)

	fmt.Println("----------")

	var u2 People
	u2.Name = "Paulo"
	u2.Age = 20
	fmt.Println(u2)

	fmt.Println("----------")

	var u3 People
	u3.Name = "Witalo"
	u3.Age = 25
	fmt.Println(u3)

	fmt.Println("----------")

	var u4 People
	u4.Name = "Carlos"
	u4.Age = 30
	fmt.Println(u4)
}