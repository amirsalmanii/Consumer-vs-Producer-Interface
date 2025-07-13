package producer

import "fmt"



// advantage 2 immediately clear where to find both interface and implementation. Player And Walkman
type Player interface {
	Play(songName string)
	// Stop() // disadvantage 2 : Limited Flexibility If you need a slightly different interface, you must modify the producer.
}

type WalkMan struct {}

func (p *WalkMan) Play(songName string) {
	fmt.Printf("Playing %s using Walkman player\n", songName)
}

func NewWalkman() Player {
	return &WalkMan{}
}
