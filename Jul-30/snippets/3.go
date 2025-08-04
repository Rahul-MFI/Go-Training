package snippets

import "fmt"

func DefaultDeclaration() {
	var s string
	var n int
	var b bool
	fmt.Printf("%q %d %t\n", s, n, b)
}
