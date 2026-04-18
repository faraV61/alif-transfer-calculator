package main

import (
	"fmt"
	"testing"
)

func ExampleMaskCard() {
	a := "1231231231231234"

	fmt.Println(MaskCard(a))

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
