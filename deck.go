package main

import "fmt"

//creating new type of 'deck' which is a slice of strings
type deck []string

//creating new function that belongs to deck type and loops through deck of cards and prints out all values
//below is a receiver function
//variable d represents the instance of the deck variable we're working with

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
