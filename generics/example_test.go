package generics

import (
	"fmt"
	"testing"

	deepcopy "github.com/barkimedes/go-deepcopy"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	id int
}

func TestGenerics(t *testing.T) {
	ints := []int64{1, 2, 3}
	floats := []float64{1.1, 2.2, 3.3}

	// fmt.Printf("sum of ints: %d, sum of floats: %.1f\n", SumIntsOrFloats[int64](ints), SumIntsOrFloats[float64](floats))
	fmt.Printf("sum of ints: %d, sum of floats: %.1f\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

// 深拷贝测试
// https://github.com/barkimedes/go-deepcopy
func TestDeepCopy(t *testing.T) {
	testArr := make([]*testStruct, 0)
	testArr = append(testArr, &testStruct{id: 1}, &testStruct{id: 2}, &testStruct{id: 3})
	copyArr, err := deepcopy.Anything(testArr)
	require.Nil(t, err)
	require.NotEmpty(t, copyArr)
	transformCopyArr := copyArr.([]*testStruct)
	require.NotEqual(t, transformCopyArr[0], testArr[0])
	// retArr := deepcopySlice(testArr)
}
