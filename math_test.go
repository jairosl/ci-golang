package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10, 10)
	
	if total != 20 {
		t.Errorf("Result invalid: result %d, expect %d", total, 20)
	}
}