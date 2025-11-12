package decoder

// Decoders records a list of functions that decode data
// of specific blocks.
type Decoders struct {
	Text         TextDecoder
	Money        MoneyDecoder
	PlayerId     PlayerIdDecoder
	PartyCount   PartyCountDecoder
	PokedexOwned PokedexOwnedDecoder
}
