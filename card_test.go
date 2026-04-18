package main

import (
	"fmt"
	"testing"
)

func ExampleMaskCard() {
	fmt.Println(MaskCard("4111 1111 1111 1111"))

}

func TestMaskCard_PanicShort(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("паиник")
		}
	}()
	MaskCard("1234")
}

func TestMaskCard_PanicLetters(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("паник")
		}
	}()
	MaskCard("abcd efhh ijkl mnop")
}

func TestMaskCard_PanicEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("паник")
		}
	}()

	MaskCard("")
}
