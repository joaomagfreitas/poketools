package poketools

// Block registers the memory region in which data is located in the save file.
type Block struct {
	Offset uint16
	Size   uint16
}

// Read reads the data which the block points to.
func (b Block) Read(data []byte) []byte {
	return data[b.Offset : b.Offset+b.Size]
}

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
