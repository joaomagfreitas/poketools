package poketools

import (
	"encoding/binary"
)

// Save identifies all data blocks that the tool can extract.
type Save struct {
	Charset  Charset
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
