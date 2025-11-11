package charset_test

import (
	"testing"

	"github.com/joaomagfreitas/poketools/charset"
)

func TestDecodable(t *testing.T) {
	testCases := []struct {
		desc      string
		charset   charset.Charset
		data      []byte
		decodable bool
	}{
		{
			desc:      "returns true if terminator is found",
			charset:   charset.Charset{Terminator: 0x50},
			data:      []byte{0x50},
			decodable: true,
		},
		{
			desc:      "returns true if no unknown character is found",
			charset:   charset.Charset{Encoding: map[byte]string{0x0: "0", 0x1: "B"}, Terminator: 0x2},
			data:      []byte{0x0, 0x1, 0x2},
			decodable: true,
		},
		{
			desc:      "returns false if an unknown character is found",
			charset:   charset.Charset{Encoding: map[byte]string{0x0: "0", 0x1: "B"}, Terminator: 0x2},
			data:      []byte{0x0, 0x1, 0x3},
			decodable: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if d := tC.charset.Decodable(tC.data); d != tC.decodable {
				t.Fail()
			}
		})
	}
}
