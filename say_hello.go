package go_say_hello

import "fmt"

func SayHello(text string) string {
	return fmt.Sprintf("Hello, %s\n", text)
}

func SumNumber(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
