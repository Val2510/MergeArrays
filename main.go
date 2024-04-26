package main

import "fmt"

func mergeArrays(arr1, arr2 []int) []int {
	merging := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merging = append(merging, arr1[i])
			i++
		} else {
			merging = append(merging, arr2[j])
			j++
		}
	}

	merging = append(merging, arr1[i:]...)
	merging = append(merging, arr2[j:]...)

	return merging
}

func main() {
	arr1 := []int{10, 30, 55, 71}
	arr2 := []int{31, 48, 50, 87, 89}

	merging := mergeArrays(arr1, arr2)
	fmt.Println(merging)
}
