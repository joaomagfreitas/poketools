package poketools

// Trainer records data about the player.
type Trainer struct {
	Name         string
	RivalName    string
	PokedexOwned []uint16
	Party        Party
	Money        uint32
	Id           uint16
}

// Party records data about the player party pokemon.
type Party struct {
	OT    []string
	Names []string
	Count uint8
}
