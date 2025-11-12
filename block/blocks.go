package block

// Blocks defines all data blocks available to be read in the save file.
type Blocks struct {
	PlayerId     Block
	Money        Block
	Name         Block
	RivalName    Block
	PartyCount   Block
	PartyOT      Block
	PartyNames   Block
	PokedexOwned Block
}
