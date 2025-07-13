package consumers

// advantage 1 : 1. Loose Coupling Each consumer defines exactly what it needs.

type AdvancedPlayer interface {
	Play(song string)
	SetVolume(level int)
}

func UseAdvancePlayer(advPlayer AdvancedPlayer, songName string, volume int) {
	advPlayer.SetVolume(volume)
	advPlayer.Play(songName)
}
