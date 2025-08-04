package snippets

import (
	"fmt"
)

type PersonFour struct {
	Name string
	Age  int
}

func Four() {
	var p1 PersonFour = PersonFour{"Rocky", 21}
	fmt.Println(p1)

	p2 := PersonFour{Name: "Rahul", Age: 21}
	fmt.Println(p2)

	p2.Age = 22
	fmt.Println(p2.Name, p2.Age)
}
