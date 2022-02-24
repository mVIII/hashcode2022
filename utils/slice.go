package utils

type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

func Map[T interface{}, K interface{}](slice []T, f func(T) K) []K {
	var result []K
	for _, e := range slice {
		result = append(result, f(e))
	}
	return result
}

func In[T comparable](arr []T, val T) bool {
	for _, elem := range arr {
		if elem == val {
			return true
		}
	}
	return false
}

func Filter[T interface{}](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, e := range slice {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

func Max[T any, K Ordered](slice []T, f func(T) K) T {
	if len(slice) == 0 {
		return *new(T)
	}
	max := slice[0]
	for _, e := range slice {
		if f(e) > f(max) {
			max = e
		}
	}
	return max
}

func Min[T any, K Ordered](slice []T, f func(T) K) T {
	if len(slice) == 0 {
		return *new(T)
	}
	min := slice[0]
	for _, e := range slice {
		if f(e) < f(min) {
			min = e
		}
	}
	return min
}

func Sort[T any, K Ordered](slice []T, f func(T) K) []T {
	sorted := make([]T, len(slice))
	_ = copy(sorted, slice)
	for i := 0; i < len(slice); i++ {
		for j := i; j > 0 && f(sorted[j-1]) > f(sorted[j]); j-- {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
		}
	}
	return sorted
}

func Find[T interface{}](arr []T, f func(T) bool) (result T) {
	r := new(T)
	for _, elem := range arr {
		if f(elem) {
			return elem
		}
	}
	return *r
}

