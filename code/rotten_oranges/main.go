package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.
*/

func main() {
	// n := 2
	// m := 2
	matrix := [][]int{{1}}
	// matrix := generateMatrix(n, m)
	// matrix := [][]int{
	// 	{1, 2},
	// 	{1, 2},
	// }
	//matrix := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	//fmt.Println(matrix)
	fmt.Println(orangesRotting(matrix))
}

type Index struct {
	i, j int
}

func orangesRotting(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}
	m := len(matrix[0])
	minutes := 0
	for {
		wg := sync.WaitGroup{}
		matrixChan := make(chan Index)
		somethingGotRotten := false
		freshOrangeExists := false
		for i := range matrix {
			for j := range matrix[i] {
				switch matrix[i][j] {
				case 1:
					if (i < n-1 && matrix[i+1][j] == 2) || (j < m-1 && matrix[i][j+1] == 2) {
						//matrix[i][j] = 2
						wg.Add(1)
						go func(i, j int) {
							matrixChan <- Index{i, j}
							wg.Done()
						}(i, j)

						somethingGotRotten = true
					} else {
						freshOrangeExists = true
					}
				case 2:
					if i < n-1 && matrix[i+1][j] == 1 {
						//matrix[i+1][j] = 2
						wg.Add(1)
						go func(i, j int) {
							matrixChan <- Index{i, j}
							wg.Done()
						}(i+1, j)

						somethingGotRotten = true
					}
					if j < m-1 && matrix[i][j+1] == 1 {
						//matrix[i][j+1] = 2
						wg.Add(1)
						go func(i, j int) {
							matrixChan <- Index{i, j}
							wg.Done()
						}(i, j+1)

						somethingGotRotten = true
					}
				}
			}
		}

		if !freshOrangeExists {
			if somethingGotRotten {
				minutes++
			}
			return minutes
		}
		minutes++
		if !somethingGotRotten {
			return -1
		}
		go func() {
			wg.Wait()
			close(matrixChan)
		}()
		//fmt.Println(matrix)
		updateMatrix(matrix, matrixChan)
		//fmt.Println(matrix)
	}
}

func updateMatrix(matrix [][]int, matrixChan chan Index) {
	for index := range matrixChan {
		matrix[index.i][index.j] = 2
	}
}

func generateMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	rand.Seed(time.Now().UnixNano())
	max := 2
	min := 0

	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			orange := rand.Intn(max-min+1) + min
			matrix[i][j] = orange
		}
	}
	return matrix
}
