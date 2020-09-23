package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck{
	cards:= deck{}
	/*
		Adding 52 cards in the slice is the goal, so instead of writing out manually we need to take an approach different
		1. A variable which contains the Suit Name for the Cards Slice.
		2. Another variable which contains the Number for those Cards
		3. We need to make nested loops where first loop is about containing the Suit names of the cards
	       and second loop is going to iterate over the card Values and finally create a new card they will append to the main
	       list of cards which is currently empty now.
		4. Finally return the card Slice
	 */
	cardSuites:= []string{"Spades","Diamonds","Hearts","clubs"}
	cardValues:=[]string{"Ace","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	for _, suit:= range cardSuites {
		for _, values:= range cardValues {
			cards = append(cards, values + "Cards of " + suit)
		}
	}
	return cards
}

/*
	Function in this case (d deck) is called a receiver which is having a print method attached to it.
	so any value of type deck can call this print method.
 */

func (d deck) print(){
	for i, card:= range d {
		fmt.Println(i,card)
	}
}

/*
	Here we are slicing the cards of type deck with number of handSize we need to take.

 */
//
//func deal(d deck, handSize int) (deck, deck) {
//	return d[:handSize], d[handSize:]
//}

/*
   Converting Type Deck --> [] string --> string --> []byte
 */

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

/*
   Saving the File, ioutil.WriteFile function returns an Error. So we are returning an error.
   We have used Receiver Function here because from main.go we can call this function.
*/

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}