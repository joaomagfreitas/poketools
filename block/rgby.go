package block

var RGBY = Blocks{
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
	RivalName: Block{
		Offset: 0x25F6,
		Size:   0xB,
	},
	PartyCount: Block{
		Offset: 0x2F2C,
		Size:   0x1,
	},
	PartyOT: Block{
		Offset: 0x2F2C + 0x110,
		Size:   0xB,
	},
	PartyNames: Block{
		Offset: 0x2F2C + 0x152,
		Size:   0xB,
	},
	PokedexOwned: Block{
		Offset: 0x25A3,
		Size:   0x13,
	},
}
