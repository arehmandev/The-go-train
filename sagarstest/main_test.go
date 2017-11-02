package main

import "testing"

var isPrimeTests = []struct {
	input   int
	want    bool
	factors int
}{
	{3, true, 2},
	{83, true, 2},
	{43, true, 2},
	{2, true, 2},
	{81, false, 5},
	{99, false, 6},
}

func TestIsPrime(t *testing.T) {
	for _, a := range isPrimeTests {
		if got := isPrime(a.input); got != a.want {
			t.Fatalf("isPrime(%d)= %t, want %t",
				a.input, got, a.want)
		}
	}
	t.Log(len(isPrimeTests), "test cases")
}

func TestGetfactors(t *testing.T) {
	for _, a := range isPrimeTests {
		if got := getfactors(a.input); got != a.factors {
			t.Fatalf("gotfactors(%d)= %d, want %d",
				a.input, got, a.factors)
		}
	}
	t.Log(len(isPrimeTests), "test cases")
}
