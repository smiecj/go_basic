package fuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	intArr   = []int{1, 2, 3, 4, 5}
	strArr   = []string{"apple", "banana", "cherry", "durian", "eggfruit"}
	floatArr = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
)

func TestReverse(t *testing.T) {
	reverse(intArr)
	reverse(strArr)

	fmt.Printf("int arr: %v\n", intArr)
	fmt.Printf("string arr: %v\n", strArr)
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"abc", "this is test string", "", "oh? my, god!"}
	for _, currentTestcase := range testcases {
		f.Add(currentTestcase)
	}

	f.Fuzz(func(t *testing.T, originStr string) {
		reverseStr := reverseString(originStr)
		reverseTwiceStr := reverseString(reverseStr)
		require.Equal(t, originStr, reverseTwiceStr)
	})
}
