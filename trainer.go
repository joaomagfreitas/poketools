package poketools

// Trainer records data about the player.
type Trainer struct {
	Name      string
	Money     uint32
	Id        uint16
	RivalName string
	Party     Party
}

// Party records data about the player party pokemon.
type Party struct {
	Count uint8
	OT    []string
	Names []string
}
