package slices

import "fmt"

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Ordered interface {
	Integer | Float | ~string
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func SliceFilter[T Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	//log(fmt.Sprintf("%s: the result is: %v", result))
	log(fmt.Sprintf("the result is: %v", result))
	return result
}
