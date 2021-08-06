package validParentheses

import (
	"testing"
)

func TestIsValid_1(t *testing.T) {
	s1 := "{}[]()"

	if isValid(s1) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}

func TestIsValid_2(t *testing.T) {
	s1 := "{[])"

	if !isValid(s1) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}

func TestIsValid_3(t *testing.T) {
	s1 := "((((()"

	if !isValid(s1) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}

func TestIsValid_4(t *testing.T) {
	s1 := "}{{}{}{}[][][]]]]]][[[["

	if !isValid(s1) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}

func TestIsValid_5(t *testing.T) {
	s1 := "{{{{[[[[(((())))]]]]}}}}"

	if isValid(s1) {
		t.Log("Accept")
	} else {
		t.Error("wrong ans")
	}
}
