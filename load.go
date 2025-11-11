package poketools

import (
	"fmt"

	"github.com/joaomagfreitas/poketools/block"
	"github.com/joaomagfreitas/poketools/charset"
	"github.com/joaomagfreitas/poketools/decoder"
)

// Load loads a Pokemon RGB/Y 32KB save file.
// Errors if save file does not match 32KB size.
func Load(data []byte) (*Save, error) {
	if l := len(data); l < 32*1024 {
		return nil, fmt.Errorf("expected save file with 32KB size (actual: %d)", l)
	}

	blk := block.RGBY
	td := [][]byte{
		blk.Name.Read(data),
		blk.RivalName.Read(data),
	}

	td = append(td, blk.PartyOT.ReadMultiple(data, 6)...)
	td = append(td, blk.PartyNames.ReadMultiple(data, 6)...)

	ok, cs := charset.Detect(
		td,
		charset.RGBY.English,
		charset.RGBY.FrenchGerman,
		charset.RGBY.ItalianSpanish,
		charset.RGBY.Japanese,
	)
	if !ok {
		return nil, fmt.Errorf("couldn't guess charset encoding of save file")
	}

	return &Save{
		Data:     data,
		Blocks:   block.RGBY,
		Decoders: decoder.RGBY(cs),
		Charset:  cs,
	}, nil
}
