package poketools

import "math"

// functions still a bit all around the place, will be refactored later.

func decodeMoney(block []byte) uint32 {
	var v int

	l := len(block)
	for _, b := range block {
		v += int(math.Pow10(2*l-2)) * int(twoDigit(b))
		l--
	}

	return uint32(v)
}

func twoDigit(b byte) uint8 {
	msb := (b >> 4) & 0xF
	lsb := b & 0xF

	return msb*10 + lsb
}
