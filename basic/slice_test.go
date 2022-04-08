// package basic go 基础包的功能验证
package basic

import (
	"reflect"
	"testing"

	"github.com/smiecj/go_common/util/log"
)

// 数组、slice 和 子slice
func TestSubSlice(t *testing.T) {
	parentSlice := make([]int, 0, 10)                          // len: 0, cap: 10
	parentSlice = append(parentSlice, []int{1, 2, 3, 4, 5}...) // len: 5, cap: 10
	sliceType := reflect.TypeOf(parentSlice)
	log.Info("[slice] slice type: %v", sliceType)

	// 子切片: 从头到中间
	// 注意: len 是元素个数，会适配成 子 slice 的元素个数，但是 cap 保持和父 slice 一致
	subSlice := parentSlice[:3] // len: 3, cap: 10
	log.Info("[sub slice] sub slice element: %v", subSlice)
	log.Info("[sub slice] sub slice len: %d, cap: %d", len(subSlice), cap(subSlice))

	// 子slice append, 非扩容
	// 子slice append，会影响子数组的 len, 但是父slice len 不变
	subSlice = append(subSlice, []int{6, 7, 8}...)                              // len: 6, cap: 10
	log.Info("[sub slice] after append, sub slice element: %v", subSlice)       // 1,2,3,6,7,8
	log.Info("[sub slice] after append, parent slice element: %v", parentSlice) // 1,2,3,6,7

	// 数组
	parentArray := [4]int{1, 2, 3, 4} // len: 4, cap: 4
	arrayType := reflect.TypeOf(parentArray)
	log.Info("[array] array type: %v", arrayType)
	log.Info("[array] array len: %d, cap: %d", len(parentArray), cap(parentArray))

	// 子数组 == 切片
	subArray := parentArray[1:3] // len: 2, cap: 3
	// 这里切片 cap 为3 原因是实际 子slice 所在空间，按照原 array 所在的空间，可在不扩容的情况下再扩
	// 一个元素，所以 cap 为3
	// https://stackoverflow.com/q/36683911
	subArrayType := reflect.TypeOf(subArray)
	log.Info("[array] sub array type: %v", subArrayType)
	log.Info("[array] sub array len: %d, cap: %d", len(subArray), cap(subArray))
}

// slice 的扩容规则
func TestSliceAppend(t *testing.T) {
	// double cap
	slice := []int{1, 2}     // len: 2, cap: 2
	slice = append(slice, 3) // len: 3, cap: 4
	log.Info("[slice] after append single elem len: %d, cap: %d", len(slice), cap(slice))

	// cap + append len
	slice = append(slice, []int{5, 6, 7, 8, 9, 10}...) // len: 9, cap: 10 (对齐)
	log.Info("[slice] after append long slice len: %d, cap: %d", len(slice), cap(slice))

	// 1.25 * cap
	longSlice := make([]int, 1024, 1024) // len: 1024, cap: 1024
	log.Info("[slice] long slice init len: %d, cap: %d", len(longSlice), cap(longSlice))
	longSlice = append(longSlice, 1) // len: 1025, cap: 1280
	log.Info("[slice] long slice after append len: %d, cap: %d", len(longSlice), cap(longSlice))
}
