package coding

import "fmt"

func Four() {
	arr := []int{1, 1, 2, 2, 2, 3, 3}
	mpp := map[int]int{}
	res := []int{}
	for i := 0; i < len(arr); i++ {
		mpp[arr[i]]++
		if mpp[arr[i]] == 1 {
			res = append(res, arr[i])
		}
	}
	fmt.Printf("The array after removing the duplicates is %d\n", res)
}
