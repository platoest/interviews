package main

import "fmt"

func main() {
	generateMatrix(6)
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	fillMatrix(1, matrix, 0, len(matrix)-1)
	fmt.Println(matrix)
	return matrix
}

func fillMatrix(startElement int, matrix [][]int, first, last int) {
	if last-first < 0 {
		return
	}
	// exit condition
	switch last - first {
	case 0:
		matrix[first][first] = startElement
		return
	case 1:
		matrix[first][first] = startElement
		matrix[first][last] = startElement + 1
		matrix[last][first] = startElement + 3
		matrix[last][last] = startElement + 2
		return
	}

	// Fill Border with elements
	newStartElement := fillBorderElements(startElement, matrix, first, last)

	// next iteration
	fillMatrix(newStartElement, matrix, first+1, last-1)
}

func fillBorderElements(startElement int, matrix [][]int, first, last int) int {
	nextElement := startElement
	for i := first; i <= last; i++ {
		matrix[first][i] = nextElement
		nextElement++
	}
	for i := first + 1; i <= last; i++ {
		matrix[i][last] = nextElement
		nextElement++
	}
	for i := last - 1; i >= first; i-- {
		matrix[last][i] = nextElement
		nextElement++
	}
	for i := last - 1; i > first; i-- {
		matrix[i][first] = nextElement
		nextElement++
	}
	return nextElement
}
