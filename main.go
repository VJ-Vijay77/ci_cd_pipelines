package main

import "fmt"

func main() {
	output := Greet()
	fmt.Println(output)
}

func Greet() string {
	return "Hello test"
}