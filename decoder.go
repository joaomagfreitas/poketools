package poketools

import "math"

// Functions still a bit all around the place, will be refactored later.

var blockName = Block{
	Offset: 0x2598,
	Size:   0xB,
}

var blockMoney = Block{
	Offset: 0x25F3,
	Size:   0x3,
}

var blockPlayerId = Block{
	Offset: 0x2605,
	Size:   0x2,
}

func decodeMoney(block []byte) uint32 {
	var v uint32

	l := len(block)
	for _, b := range block {
		v += uint32(math.Pow10(2*l-2)) * uint32(twoDigit(b))
		l--
	}

	return v
}

func twoDigit(b byte) uint8 {
	msb := (b >> 4) & 0xF
	lsb := b & 0xF

	return msb*10 + lsb
}
