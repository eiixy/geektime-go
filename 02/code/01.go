package code

import (
	"errors"
	"fmt"
)

func Positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}

func Check(n int) {
	pos, err := Positive(n)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

// errors are value
func Test() {
	if ErrNamedType == New("EOF") {
		fmt.Println("named type error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}
