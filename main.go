package main

import "fmt"

func main() {
	result, error := hello("helloword")
	if error != nil {
		fmt.Println("error")
	}
	fmt.Println(result)
}

func hello(s string) (string, error) {
	return s, nil
}
