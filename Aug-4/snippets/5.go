package snippets

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person) {
	p.Age += 1
}

func Five() {
	p := Person{"Rocky", 22}
	updateAge(&p)
	fmt.Println(p.Age)
}
