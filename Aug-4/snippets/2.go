package snippets

import (
	"fmt"
)

func update(n *int) {
	*n = 2
}

func Two() {
	x := 10
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}
