package mypack

import "testing"

func TestValidip(t *testing.T) {

	ret := Validip("192.168.043.005")
	if ret != 1 {
		t.Errorf("Correct check not done")
	}
}

func TestNonalpha(t *testing.T) {

	ret := Nonalpha("sa;[];")
	if ret != "sa" {
		t.Errorf("wrong output")
	}
}
