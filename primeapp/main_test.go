package main

import "testing"

func Test_isPrime(t *testing.T) {
	// n := 0
	// result, msg := isPrime(n)
	// if result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false, message: '%s'", n, msg)
	// }
	primeTests := []struct {
		name string	
		testNum int
		expected bool
	}{
		{ name: "Negative", testNum: -1, expected: false, },
		{ name: "One", testNum: 1, expected: false, },
		{ name: "Prime", testNum: 7, expected: true, },
		{ name: "Not prime", testNum: 6, expected: false, },
	}

	for _, testCase := range primeTests {
		result, msg := isPrime(testCase.testNum)
		if result != testCase.expected {
			t.Errorf(
				"%s: Test case for number %d failed, expected %t, but got %t. Message from function: '%s'",
				testCase.name,
				testCase.testNum,
				testCase.expected,
				result,
				msg,
			)
		}
	}
}