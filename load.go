package poketools

import (
	"fmt"
)

// Load loads a Pokemon RGB/Y 32KB save file.
// Errors if save file does not match 32KB size.
func Load(data []byte) (*Save, error) {
	if l := len(data); l != 32*1024 {
		return nil, fmt.Errorf("expected save file with 32KB size (actual: %d)", l)
	}

	return &Save{
		Data:     data,
		PlayerId: blockPlayerId,
		Money:    blockMoney,
		Name:     blockName,
		Charset:  EnglishCharset,
	}, nil
}
