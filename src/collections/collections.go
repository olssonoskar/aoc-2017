package collections

import (
	"cmp"
	"slices"
)

func SortAsc[T, E cmp.Ordered](m map[T]E) []E {
	values := make([]E, 0)
	for _, val := range m {
		values = append(values, val)
	}
	slices.SortFunc(values, cmp.Compare)
	return values
}

func SortDesc[T, E cmp.Ordered](m map[T]E) []E {
	values := SortAsc(m)
	slices.Reverse(values)
	return values
}
