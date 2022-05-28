package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("The Length of deck is not 16, it is %v",len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Not the Ace of Spades, %v instead",d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Not the Four of Clubs, %v instead",d[len(d)-1])
	}
}

func TestSaveFileAndReadFile(t *testing.T) {
	os.Remove("_testingdeck")

	deck := newDeck()
	deck.saveFile("_testingdeck")

	loadedDeck := readFile("_testingdeck")

	if len(loadedDeck) != 16 {
		t.Errorf("Not expected slice length, got %v instead", len(loadedDeck))
	}

	os.Remove("_testingdeck")
}