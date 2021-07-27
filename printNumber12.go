package main

import "fmt"

func printFinalNumber(num []byte)  {
	beginIndex := 0
	for i, b := range num {
		if b != '0' {
			beginIndex = i
			break
		}
	}

	fmt.Println(string(num[beginIndex:]))
}

func printNumberRecursively(num []byte, index int)  {
	if index == len(num) {
		printFinalNumber(num)
		return
	}

	for i := 0; i < 10; i++ {
		num[index] = byte(i + '0')
		printNumberRecursively(num, index+1)
	}
}

func PrintNumber(n int) {
	if n <= 0 {
		return
	}

	num := make([]byte, n)
	printNumberRecursively(num, 0)
}