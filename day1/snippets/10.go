package snippets

import "fmt"

func calc() (int, int) {
	return 1, 1 // comment 1,1
}
func ReturnWithoutValue() {
	x, y := calc()
	fmt.Print(x, y)
}
