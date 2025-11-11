package decoder_test

import (
	"testing"

	"github.com/joaomagfreitas/poketools/charset"
	"github.com/joaomagfreitas/poketools/decoder"
)

func TestRGBYMoneyDecoder(t *testing.T) {
	ds := decoder.RGBY(charset.Charset{})
	td := map[uint32][]byte{
		0:      {0x0, 0x0, 0x0},
		1:      {0x0, 0x0, 0x1},
		9:      {0x0, 0x0, 0x9},
		12:     {0x0, 0x0, 0x12},
		25:     {0x0, 0x0, 0x25},
		101:    {0x0, 0x1, 0x01},
		255:    {0x0, 0x2, 0x55},
		3455:   {0x0, 0x34, 0x55},
		583210: {0x58, 0x32, 0x10},
	}

	for tm, d := range td {
		if m := ds.Money(d); m != tm {
			t.Fatalf("%d != %d", tm, m)
		}
	}
}

func TestRGBYPartyCountDecoder(t *testing.T) {
	ds := decoder.RGBY(charset.Charset{})
	td := map[uint8][]byte{
		1: {0x1},
		2: {0x2},
		3: {0x3},
		6: {0x6},
	}

	for tpc, d := range td {
		if pc := ds.PartyCount(d); pc != tpc {
			t.Fatalf("%d != %d", tpc, pc)
		}
	}
}
