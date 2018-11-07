package main

import "testing"

func TestCheckPassword(t *testing.T) {
	testSecret := "testSecret"
	testSecretHashed := HashPassword(testSecret)
	pass, err := CheckPassword(testSecret, testSecretHashed)
	if err != nil {
		t.Fatal(err)
	}
	if !pass {
		t.Fatal("CheckPassword == false, want true")
	}
}

func BenchmarkCheckPassword(b *testing.B) {
	// run the Fib function b.N times
	testSecret := "testSecret"
	testSecretHashed := HashPassword(testSecret)
	for n := 0; n < b.N; n++ {
		if _, err := CheckPassword(testSecret, testSecretHashed); err != nil {
			b.Fatal(err)
		}
	}
}
