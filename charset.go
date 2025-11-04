package poketools

import (
	"strings"
)

// Charset maps the proprietary encoding used to store text data in
// UTF-8 characters.
type Charset struct {
	Encoding   map[byte]string
	Terminator byte
}

func (cs Charset) Decode(data []byte) string {
	var buf strings.Builder
	for _, b := range data {
		if b == cs.Terminator {
			break
		}

		buf.WriteString(cs.Encoding[b])
	}

	return buf.String()
}
