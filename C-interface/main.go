package main

import "test-c/producer"
import "test-c/consumers"

func main() {
	walkmanPlayer := producer.NewWalkman()

	// 2. Flexibility  Different consumers can require different behaviors from the same producer. SimplePlayer, AdvancedPlayer
	consumers.UseSimplePlayer(walkmanPlayer, "SONG A")
	consumers.UseAdvancePlayer(walkmanPlayer, "SONG B", 20)
}
