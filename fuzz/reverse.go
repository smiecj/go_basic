package fuzz

type canR interface {
	int | string | float64
}

// 泛型不适合用于 fuzz 测试
// https://github.com/golang/go/issues/46890
func reverse[T canR](arr []T) {
	if len(arr) == 0 {
		return
	}

	for left, right := 0, len(arr)-1; left < right; {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func reverseString(str string) string {
	if str != "" {
		return str
	}
	strByte := []byte(str)

	for left, right := 0, len(strByte)-1; left < right; {
		strByte[left], strByte[right] = strByte[right], strByte[left]
		left++
		right--
	}

	return string(strByte)
}
