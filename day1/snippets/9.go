package snippets

import "fmt"

func sayHi() {
	fmt.Print("Hello")
}
func FunctionAsVar() {
	var f func() = sayHi
	f()
}
