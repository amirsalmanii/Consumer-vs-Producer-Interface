package producer

import "fmt"

// Note: No interfaces defined here, just concrete implementations

type Walkman struct{}

func (p *Walkman) Play(songName string) {
	fmt.Printf("Playing %s using Walkman player\n", songName)
}

func (p *Walkman) SetVolume(volume int) {
	fmt.Printf("Set volume to %d ", volume)
}

func NewWalkman() *Walkman {
	return &Walkman{}
}
