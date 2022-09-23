package main

import (
	"fmt"
)

// Даны три неубывающих массива чисел. Найти число, которое присутствует во всех трех массивах.
// Input: [1,2,4,5], [3,3,4], [2,3,4,5,6]
// Output: 4
func main() {
	indexes := []int{0, 0, 0}
	arrays := [][]int{
		{1, 2, 4, 5, 7},
		{3, 3, 7, 7},
		{2, 3, 4, 5, 6, 7},
	}
	for {
		for i := 0; i < len(arrays); i++ {
			if indexes[i]+1 > len(arrays[i]) {
				fmt.Printf("array %d doesn't have any values left", i)
				return
			}
		}
		val0 := arrays[0][indexes[0]]
		val1 := arrays[1][indexes[1]]
		val2 := arrays[1][indexes[2]]
		if val0 == val1 {
			if val1 == val2 {
				fmt.Printf("%d", val0)
				return
			} else if val1 > val2 {
				indexes[0]++
				indexes[1]++
			} else {
				indexes[2]++
			}
			continue
		} else if val1 == val2 {
			if val0 > val2 {
				indexes[1]++
				indexes[2]++
			} else {
				indexes[0]++
			}
		} else if val0 < val1 {
			if val0 < val2 {
				indexes[0]++
			} else {
				indexes[2]++
			}
		} else if val1 < val2 {
			indexes[1]++
		} else {
			indexes[2]++
		}
	}
}
