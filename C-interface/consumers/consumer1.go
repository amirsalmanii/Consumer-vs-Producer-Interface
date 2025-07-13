package consumers

// advantage 1 : 1. Loose Coupling Each consumer defines exactly what it needs.

type SimplePlayer interface {
	Play(songName string)
}

func UseSimplePlayer(simplePlayer SimplePlayer, songName string) {
	simplePlayer.Play(songName)
}
