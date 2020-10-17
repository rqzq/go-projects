package main

import (
	"os"
	"testing"
)

func TestNewDeck( t *testing.T){
	d:= newDeck()
	if len(d)!= 40{
		t.Errorf("Expected value 40, but got %v", len(d))
	}
	if d[0] != "Ace Cards of Spades"{
		t.Errorf("Expected Ace of Spades, but got %v",d[0])
	}
	if d[len(d)-1] != "Nine Cards of clubs"{
		t.Errorf("Exepected Nine of clubs, but got %v",d[len(d)-1])
	}
}

func TestSavetoFileAndNewDeckfromFile( t *testing.T){
	os.Remove("_deckTesting")
	d:= newDeck()
	d.saveToFile("_deckTesting")
	loadedDeck:= newDeckFromFile("_deckTesting")
	if len(loadedDeck) !=40{
		t.Errorf("Expected Deck size of 40, but got %v",len(loadedDeck))
	}
	os.Remove("_deckTesting")

}
