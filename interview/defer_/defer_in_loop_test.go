package defer_

import (
	"fmt"
	"testing"
)

func TestDeferInForLoopWrong(t *testing.T) {
	for index := 0; index < 5; index++ {
		defer func() {
			fmt.Println(index)
		}()
	}

	fmt.Println("defer loop")
}

func TestDeferInForLoopRight(t *testing.T) {
	for index := 0; index < 5; index++ {
		defer func(index int) {
			fmt.Println(index)
		}(index)
	}

	fmt.Println("defer loop")
}

func TestDeferAndReturnAnonymous(t *testing.T) {
	f := func() int {
		a := 1
		defer func() {
			a = 3
		}()
		a = 2
		return a
	}

	fmt.Println(f())
}

func TestDeferAndReturnVarName(t *testing.T) {
	f := func() (a int) {
		a = 1
		defer func() {
			a = 3
		}()
		a = 2
		return
	}

	fmt.Println(f())
}

func TestDeferPassFuncGenParam(t *testing.T) {
	s := "a"

	defer fmt.Println(func() string {
		return s
	}())

	s = "b"
}

func TestDeferAndPanic(t *testing.T) {
	defer func() {
		r := recover()
		fmt.Println("recovered:", r)
	}()

	panic("error")
}
