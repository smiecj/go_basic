package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

var (
	byteLen   = 1024
	longStr   = strings.Repeat("S", byteLen)
	longBytes = bytes.Repeat([]byte{'S'}, byteLen)
)

// 强转写法
func fcBytesToString(bytes []byte) string {
	return string(bytes)
}

func fcStringToBytes(str string) []byte {
	return []byte(str)
}

// func byteSliceToString(bytes []byte) string {
// 	var s string
// 	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
// 	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	stringHeader.Data = sliceHeader.Data
// 	stringHeader.Len = sliceHeader.Len
// 	return s
// }

func byteSliceToString(bytes []byte) string {
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

// func stringToByteSlice(s string) (bytes []byte) {
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
// 	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh.Data = sh.Data
// 	bh.Len = sh.Len
// 	bh.Cap = sh.Len
// 	return
// }

func stringToByteSlice(s string) (bytes []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func TestConvert(t *testing.T) {
	fmt.Println(byteSliceToString([]byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}))
	fmt.Printf("%v\n", stringToByteSlice("hello, world"))
}

func BenchmarkConvertBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteSliceToString(longBytes)
	}
}

func BenchmarkConvertStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringToByteSlice(longStr)
	}
}

func BenchmarkForceConvertBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fcBytesToString(longBytes)
	}
}

func BenchmarkForceConvertStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fcStringToBytes(longStr)
	}
}
