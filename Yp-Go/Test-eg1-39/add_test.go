package main

import "testing"

func TestAddPositive(t *testing.T) {

	sum, err := Add(1, 2)

	if err != nil {
		t.Error("unexpected error")
	}

	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {

	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}

	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expted error not be nil")
	}

	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expted error not be nil")
	}

}
