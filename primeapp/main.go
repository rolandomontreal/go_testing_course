package main

import "fmt"

func main() {
	n := 2

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		// msg := fmt.Sprintf("%d", n)
		return false, fmt.Sprintf("%d is not prime by definition!", n)
	}
	
	// Negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n / 2; i++ {
		// not a prime number
		if n % i == 0 {
			return false, fmt.Sprintf("%d is not prime, it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}