package main

import "testing"

func Test_isPrime(t *testing.T) {
	n := 0
	result, msg := isPrime(n)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false, message: '%s'", n, msg)
	}
}