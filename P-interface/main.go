// here is consumer side

package main

import "test-p/producer" // disadvantages 1  Tight coupling to the producer package

func main() {
	walkManPlayer := producer.NewWalkman()

	// advantage 1 : // Consumer can't access internal details, only the interface
	walkManPlayer.Play("SONG A")
}