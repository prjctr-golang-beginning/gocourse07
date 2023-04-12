package main

import (
	"fmt"
)

func main() {
	input := "Hello, world!" // вхідний рядок
	c1 := make(chan string)  // канал для першої горутини
	c2 := make(chan string)  // канал для другої горутини

	// перша горутина - додає "Hello" до рядка
	go func() {
		c1 <- "Hello"
	}()

	// друга горутина - додає ", world!" до рядка
	go func() {
		c2 <- ", world!"
	}()

	// об'єднуємо канали і виводимо результат
	output := input + <-c1 + <-c2
	fmt.Println(output)
}
