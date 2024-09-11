package slicefunctions

import (
	"fmt"
	"slices"
)

func TestSliceFunctions() {
	someSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Repeat(someSlice, 2))
	slices.Sort(someSlice)
	slices.Reverse(someSlice)
	someSlice2 := slices.Clone(someSlice)
	someSlice2[0] = 0
	fmt.Println(someSlice, someSlice2)
	fmt.Println(slices.Concat([]int{1, 2, 3}, []int{4, 5}, []int{7, 8}))
	if slices.Contains([]int{1, 2, 3}, 3) {
		fmt.Println("yes, it contains 3")
	}
	fmt.Println(slices.Delete([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, 4))
	if slices.Equal([]int{1, 2}, []int{1, 2}) {
		fmt.Println("They have the same value")
	}
	if slices.IsSorted([]int{1, 2, 3}) {
		fmt.Println("Sorted")
	}
	fmt.Println("Max value: ", slices.Max(someSlice), " Min Value: ", slices.Min(someSlice))
	fmt.Println(slices.Replace([]int{1, 2, 3, 4, 5}, 1, 2, []int{6, 7, 8}...))
}
