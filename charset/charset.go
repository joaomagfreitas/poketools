package charset

import "strings"

// Charset maps the proprietary encoding used to store text data in
// UTF-8 characters.
type Charset struct {
	Encoding   map[byte]string
	Terminator byte
}

// Decode decodes data block that is encoded using the charset.
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

// Decodable checks whether a data block can be decoded using this charset.
func (cs Charset) Decodable(data []byte) bool {
	for _, b := range data {
		if b == cs.Terminator {
			break
		}

		if _, ok := cs.Encoding[b]; !ok {
			return false
		}
	}

	return true
}
