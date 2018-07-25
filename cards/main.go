package main

import (
	"fmt"
)

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" // := is read by go compiler := is only used when defining a new variable
	card = "Five of Diamonds"
	fmt.Println(card)
}
