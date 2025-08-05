package coding

import "fmt"

func Six() {
	arr := []int{2, 7, 11, 15}
	target := 13
	res := []int{-1, -1}
	i, j := 0, len(arr)-1
	for i <= j {
		curr := arr[i] + arr[j]
		if curr > target {
			j--
		} else if curr == target {
			res[0] = i
			res[1] = j
			break
		} else {
			i++
		}
	}
	fmt.Printf("The index of the target sum is %d\n", res)
}
