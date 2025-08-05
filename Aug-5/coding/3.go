package coding

import "fmt"

func mergeSlices(a, b []int) []int {
	res := []int{}
	i, j := 0, 0
	n, m := len(a), len(b)
	for i < n && j < m {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

func Three() {
	slice1 := []int{1, 3, 5}
	slice2 := []int{2, 4, 6, 7}
	res := mergeSlices(slice1, slice2)
	fmt.Printf("The merged slice is %d\n", res)
}
