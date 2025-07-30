package snippets

import "fmt"

func test() (a int, b int) {
	a = 1
	b = 2
	return
}
func NameReturnValue() {
	x, y := test()
	fmt.Println(x, y)
}
