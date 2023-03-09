package main

/*
User interface contains NumericInput control, which accepts only digits.

Extend NumericInput structure so that:

    It implements UserInput interface.
    Add(rune) should add only decimal digits to the input. Other runes should be ignored.
    GetValue() should return the current input.

For example, the following code should output "10":

var input UserInput = &NumericInput{}
input.Add('1')
input.Add('a')
input.Add('0')
fmt.Println(input.GetValue())

Reference - https://www.testdome.com/tests/golang-online-test/123
*/
import (
	"fmt"
	"unicode"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (u *NumericInput) Add(r rune) {
	if unicode.IsDigit(r) {
		u.input += string(r)
	}
}

func (u NumericInput) GetValue() string {
	return u.input
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
