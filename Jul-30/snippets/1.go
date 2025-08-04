package snippets

import "fmt"

var x = 10

func Shadowing() {
	fmt.Println(x)
	x := 20
	fmt.Println(x)
}
