package charset_test

import (
	"reflect"
	"testing"

	"github.com/joaomagfreitas/poketools/charset"
)

func TestDetect(t *testing.T) {
	testCases := []struct {
		desc     string
		charset  charset.Charset
		data     [][]byte
		charsets []charset.Charset
		ok       bool
	}{
		{
			desc:     "returns no charset if no data blocks are provided",
			charsets: []charset.Charset{{}},
		},
		{
			desc: "returns no charset if no charsets are provided",
			data: [][]byte{{}},
		},
		{
			desc: "returns no charset if none of provided charsets can decode data blocks",
			charsets: []charset.Charset{
				{Encoding: map[byte]string{0x3: "3", 0x4: "4"}, Terminator: 0x5},
				{Encoding: map[byte]string{0x5: "5", 0x6: "6"}, Terminator: 0x7},
			},
			data: [][]byte{{0x0, 0x1, 0x2}},
		},
		{
			desc: "returns charset with most decodable data blocks ",
			charsets: []charset.Charset{
				{Encoding: map[byte]string{0x3: "3", 0x4: "4"}, Terminator: 0x5},
				{Encoding: map[byte]string{0x5: "5", 0x6: "6"}, Terminator: 0x7},
			},
			data:    [][]byte{{0x3}, {0x6}, {0x7}},
			ok:      true,
			charset: charset.Charset{Encoding: map[byte]string{0x5: "5", 0x6: "6"}, Terminator: 0x7},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if ok, cs := charset.Detect(tC.data, tC.charsets...); ok != tC.ok || !reflect.DeepEqual(cs, tC.charset) {
				t.Fail()
			}
		})
	}
}
