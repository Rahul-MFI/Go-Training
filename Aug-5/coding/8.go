package coding

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Eight() {
	num := 13
	fmt.Printf("Is %d a prime number? %v\n", num, isPrime(num))
}
