package decoder

import (
	"encoding/binary"
	"math"

	"github.com/joaomagfreitas/poketools/charset"
)

func RGBY(cs charset.Charset) Decoders {
	return Decoders{
		Text:         cs.Decode,
		Money:        rgbyDecodeMoney,
		PlayerId:     binary.BigEndian.Uint16,
		PartyCount:   rgbyDecodePartyCount,
		PokedexOwned: rgbyDecodePokedexOwned,
	}
}

func rgbyDecodePokedexOwned(data []byte) []uint16 {
	owned := []uint16{}
	for i := range uint16(151) {
		if data[i>>3]>>(i&7)&1 == 0 {
			continue
		}

		owned = append(owned, i+1)
	}

	return owned
}

func rgbyDecodePartyCount(data []byte) uint8 {
	return data[0]
}

func rgbyDecodeMoney(data []byte) uint32 {
	var v uint32

	l := len(data)
	for _, b := range data {
		v += uint32(math.Pow10(2*l-2)) * uint32(rgbyTwoDigit(b))
		l--
	}

	return v
}

func rgbyTwoDigit(b byte) uint8 {
	msb := (b >> 4) & 0xF
	lsb := b & 0xF

	return msb*10 + lsb
}
