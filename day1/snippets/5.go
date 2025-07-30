package snippets

import "fmt"

func sum(nums ...int) int {
	tot := 0
	for _, i := range nums {
		tot += i
	}
	return tot
}
func VariadicFun() {
	data := []int{1, 2, 3}
	fmt.Println(sum(data...))
}
