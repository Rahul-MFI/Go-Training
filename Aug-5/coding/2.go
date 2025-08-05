package coding

import "fmt"

func swap(arr []int, n int) []int {
	st := arr[0]
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = st
	return arr
}

func Two() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	n := len(arr)
	kk := k % n
	for i := 0; i < kk; i++ {
		arr = swap(arr, n)
	}
	fmt.Printf("The rotated array upto %d times is %d\n", k, arr)
}
