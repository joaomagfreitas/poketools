package block

// Block registers the memory region in which data is located in the save file.
type Block struct {
	Offset uint16
	Size   uint16
}

// Read reads the data which the block points to.
func (b Block) Read(data []byte) []byte {
	return data[b.Offset : b.Offset+b.Size]
}

// ReadMultiple reads the data of [count] sequentially divided blocks.
func (b Block) ReadMultiple(data []byte, count int) [][]byte {
	d := make([][]byte, count)
	for i := range count {
		off := b.Offset + (b.Size * uint16(i))
		d[i] = data[off : off+b.Size]
	}

	return d
}
