package coding

import "fmt"

func Five() {
	s := "gogopher"
	mpp := map[string]int{}
	for i := 0; i < len(s); i++ {
		mpp[string(s[i])]++
	}
	fmt.Printf("The count of characters are %+v\n", mpp)
}
