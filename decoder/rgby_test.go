package decoder_test

import (
	"slices"
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

func TestRGBYPokedexOwnedDecoder(t *testing.T) {
	ds := decoder.RGBY(charset.Charset{})

	td := []byte{
		0x4B, // 0100 1011 (1, 2, 4 and 7 pokemon)
		0x0,
		0x89, // 1000 1001 (17, 20, and 24 pokemon)
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}

	tpo := []uint16{1, 2, 4, 7, 17, 20, 24}
	if po := ds.PokedexOwned(td); !slices.Equal(po, tpo) {
		t.Fatalf("%d != %d", tpo, po)
	}
}
