package main

import (
	"fmt"
	"strconv"
)

func main() {
	card := "Ace of spades"
	number := 14
	t := strconv.Itoa(number)
	fmt.Println("Hello World! " + card + " " + t)
}
