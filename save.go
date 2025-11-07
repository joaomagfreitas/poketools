package poketools

import (
	"encoding/binary"

	"github.com/joaomagfreitas/poketools/block"
	"github.com/joaomagfreitas/poketools/charset"
)

// Save identifies all data blocks that the tool can extract.
type Save struct {
	Charset charset.Charset
	Data    []byte
	Blocks  block.Blocks
}

// Reads trainer data from save file.
func (s Save) Trainer() Trainer {
	blk := s.Blocks
	n := blk.Name.Read(s.Data)
	id := blk.PlayerId.Read(s.Data)
	m := blk.Money.Read(s.Data)

	return Trainer{
		Id:    binary.BigEndian.Uint16(id),
		Money: decodeMoney(m),
		Name:  s.Charset.Decode(n),
	}
}
