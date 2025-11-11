package decoder

// Alias for decoding text data.
type TextDecoder = func(data []byte) string

// Alias for decoding money data.
type MoneyDecoder = func(data []byte) uint32

// Alias for decoding player id data.
type PlayerIdDecoder = func(data []byte) uint16

// Alias for decoding party pokemon count.
type PartyCountDecoder = func(data []byte) uint8
