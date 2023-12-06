package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	number := 13
	if isPrime(number) {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}
}
