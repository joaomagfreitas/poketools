package poketools

import (
	"encoding/binary"

	"github.com/joaomagfreitas/poketools/charset"
)

// Save identifies all data blocks that the tool can extract.
type Save struct {
	Charset  charset.Charset
	Data     []byte
	PlayerId Block
	Money    Block
	Name     Block
}

// Reads trainer data from save file.
func (s Save) Trainer() Trainer {
	n := s.Name.Read(s.Data)
	id := s.PlayerId.Read(s.Data)
	m := s.Money.Read(s.Data)

	return Trainer{
		Id:    binary.BigEndian.Uint16(id),
		Money: decodeMoney(m),
		Name:  s.Charset.Decode(n),
	}
}
