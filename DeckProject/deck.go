package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck{
	cards:= deck{}
	cardSuites:= []string{"Spades","Diamonds","Hearts","clubs"}
	cardValues:=[]string{"Ace","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	for _, suit:= range cardSuites {
		for _, values:= range cardValues {
			cards = append(cards, values + "Cards of " + suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card:= range d {
		fmt.Println(i,card)
	}
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck{
	bs, err:= ioutil.ReadFile(filename)
	if err!= nil{
		fmt.Println("Error",err)
		os.Exit(503)
	}
	s:= strings.Split(string(bs),",")
	return deck(s)
}