package main

import "fmt"

func main() {
	output := Greet()
	fmt.Println(output)
	AddingTwoNumbers := Add(4,5)
	fmt.Println(AddingTwoNumbers)

	SubstractingTwoNumbers := Substract(5,2)
	fmt.Println(SubstractingTwoNumbers)

	AddTailToString := Tail("sample")
	fmt.Println(AddTailToString)
}

func Tail(s string) string {
	t := "-TailAdded"
	new := s + t
	return new
}


func Add(a int,b int) int{
	sum := a + b
	return sum
}
func Substract(a int,b int) int{
	sum := a - b
	return sum
}

func Greet() string {
	return "Hello test"
}