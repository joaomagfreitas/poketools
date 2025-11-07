package block

var RGBY = rgbyBlocks{
	All: Blocks{
		PlayerId: Block{
			Offset: 0x2605,
			Size:   0x2,
		},
		Money: Block{
			Offset: 0x25F3,
			Size:   0x3,
		},
		Name: Block{
			Offset: 0x2598,
			Size:   0xB,
		},
	},
}

type rgbyBlocks struct {
	// afaik, gen 1 games use the same save data structure layout
	// for north america, europe and japanese titles
	All Blocks
}
