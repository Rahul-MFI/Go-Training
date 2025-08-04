package snippets

import (
	"fmt"
)

func Three() {
	x := 10
	var p *int = &x
	var pp **int = &p

	fmt.Println("x:", x)
	fmt.Println("*p:", *p)
	fmt.Println("**p:", **pp)
	**pp = 20
	fmt.Println("Updated x:", x)
}
