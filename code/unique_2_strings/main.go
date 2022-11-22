package main

import (
	"fmt"
	"strings"
)

/*
	Write a simple method to print out the unique characters across two given strings
	Example
	S1 = Los Angeleso
	S2 = Las Vegas
	Output is = O, V, N
	Solutions needs to be in O(n) Time & Space Complexity
*/

func uniqueChar(s1, s2 string) {
	chars := strings.Split(s1, "")
	m := createMap(s2)
	unique := make(map[string]bool)
	for _, c := range chars {
		_, ok := m[c]
		if !ok {
			unique[c] = true
		}
	}
	chars = strings.Split(s2, "")
	m = createMap(s1)
	for _, c := range chars {
		_, ok := m[c]
		if !ok {
			unique[c] = true
		}
	}

	for c := range unique {
		fmt.Println(c)
	}
}

func createMap(s string) map[string]bool {
	m := make(map[string]bool)
	chars := strings.Split(s, "")
	for _, c := range chars {
		m[c] = true
	}
	return m
}

func main() {
	s1 := strings.ToLower("Los Angeleso")
	s2 := strings.ToLower("Las Vegas")

	uniqueChar(s1, s2)
}
