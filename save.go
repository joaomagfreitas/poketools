package poketools

import (
	"github.com/joaomagfreitas/poketools/block"
	"github.com/joaomagfreitas/poketools/charset"
	"github.com/joaomagfreitas/poketools/decoder"
	"github.com/joaomagfreitas/stdx/slicesx"
)

// Save identifies all data blocks that the tool can extract.
type Save struct {
	Decoders decoder.Decoders
	Charset  charset.Charset
	Data     []byte
	Blocks   block.Blocks
}

// Reads trainer data from save file.
func (s Save) Trainer() Trainer {
	blk := s.Blocks
	pc := s.Decoders.PartyCount(blk.PartyCount.Read(s.Data))

	return Trainer{
		Id:        s.Decoders.PlayerId(blk.PlayerId.Read(s.Data)),
		Money:     s.Decoders.Money(blk.Money.Read(s.Data)),
		Name:      s.Decoders.Text(blk.Name.Read(s.Data)),
		RivalName: s.Decoders.Text(blk.RivalName.Read(s.Data)),
		Party: Party{
			Count: pc,
			OT:    slicesx.Map(blk.PartyOT.ReadMultiple(s.Data, uint16(pc)), s.Decoders.Text),
			Names: slicesx.Map(blk.PartyNames.ReadMultiple(s.Data, uint16(pc)), s.Decoders.Text),
		},
		PokedexOwned: s.Decoders.PokedexOwned(blk.PokedexOwned.Read(s.Data)),
	}
}
