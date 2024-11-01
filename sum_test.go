package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("O resultado é diferente de 5")
	}
}

func TestMain(t *testing.T) {
	main()
}
