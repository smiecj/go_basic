package generics

// SumInts adds together the values of m.
func SumInts(m []int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m []float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// func SumIntsOrFloats[V int64 | float64](m []V) V {
// 	var s V
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

type Number interface {
	int64 | float64
}

// type T[T any] struct{}
// func (T[T]) m() {}

func SumIntsOrFloats[V Number](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 自己实现深拷贝（数组类型不支持设置成泛型）
// https://stackoverflow.com/a/54973735
type T interface {
	any
}

func deepcopySlice(source []*T) (ret []*T) {
	ret = make([]*T, len(source))
	for index, pointer := range source {
		if pointer == nil {
			continue
		}
		value := *pointer
		ret[index] = &value
	}
	return
}

/* func deepcopySlice(source []*T) (ret []*T) {
	ret = make([]*T, len(source))
	for index, pointer := range source {
		if pointer == nil {
			continue
		}
		value := *pointer
		ret[index] = &value
	}
	return
}
*/
