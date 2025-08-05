package coding

import "fmt"

func One() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	fmt.Printf("The Reversed array is %d\n", arr)
}
