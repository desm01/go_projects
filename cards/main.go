package main

func main() {
	c := newDeck()
	c.shuffle()
	c.saveToFile("my_cards")
}
