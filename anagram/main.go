package main

import (
	"fmt"
)

func MergeSort(data []int32) []int32 {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])

	return merge(left, right)
}
func merge(left, right []int32) []int32 {
	result := make([]int32, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

func anagram(list []string) [][]string {
	var temp = [][]string{}
	var convert []rune
	var sortingText []int32
	var result = make(map[string][]string)

	for _, v := range list {
		convert = []rune(v)
		sortingText = MergeSort(convert)
		result[string(sortingText)] = append(result[string(sortingText)], v)
	}

	for _, el := range result {
		temp = append(temp, el)
	}
	return temp
}

func main() {
	fmt.Println(anagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}
