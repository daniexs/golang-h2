package main

import "fmt"

func main() {
	result := AlayGen("hello", "word", "zzzz")
	fmt.Println(result)

	result2 := soal2(8)
	fmt.Println(result2)
}

func AlayGen(value ...string) string {
	result := ""
	for _, v := range value {
		result = result + v + " "
	}

	temp := ""
	for i := 0; i < len(result); i++ {
		word := modifyWord(string(result[i]))
		temp = temp + word
	}

	return temp
}

func modifyWord(value string) string {
	switch value {
	case "a":
		return "4"
	case "e":
		return "3"
	case "i":
		return "!"
	case "l":
		return "1"
	case "n":
		return "N"
	case "s":
		return "5"
	case "x":
		return "*"
	default:
		return value
	}
}

func soal2(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
