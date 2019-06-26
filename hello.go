package main

import (
    "strconv"
    "fmt"
)

func main() {
	card := "Ace of spades"
	number := 14
	t := strconv.Itoa(number)
	fmt.Println("Hello World! " + card + " " + t)
}
