package main

import "fmt"

/*
1 -> A
2 - > B

25 ->Y
26 -> Z

27 -> AA

52 ->  AZ
53 -> BA

78 -> BZ
79 -> CA

702 -> ZZ
703 -> AAA

728 -> AAZ
*/

func main() {
	number := 728
	//number := 79
	leftover := (number - 1) % 26
	other := number - leftover + 1
	if other > 0 {
		printColumnLetter(other)
	}
	fmt.Print(string(rune('A' + leftover)))
}

func printColumnLetter(number int) {
	number = number / 26
	if number > 26 {
		printColumnLetter(number)
	}
	leftover := (number - 1) % 26
	fmt.Print(string(rune('A' + leftover)))
}
