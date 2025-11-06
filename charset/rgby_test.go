package charset_test

import (
	"testing"

	"github.com/joaomagfreitas/poketools/charset"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		desc    string
		decoded string
		charset charset.Charset
		encoded []byte
	}{
		{
			desc:    "english charset",
			encoded: []byte{128, 146, 135, 80, 137, 128, 130, 138, 80, 141, 132},
			decoded: "ASH",
			charset: charset.RGBY.English,
		},
		{
			desc:    "french/german charset",
			encoded: []byte{128, 219, 202},
			decoded: "As'û",
			charset: charset.RGBY.FrenchGerman,
		},
		{
			desc:    "italian/spanish charset",
			encoded: []byte{136, 240, 140},
			decoded: "I$M",
			charset: charset.RGBY.ItalianSpanish,
		},
		{
			desc:    "japanese charset",
			encoded: []byte{20, 40, 60},
			decoded: "ナ゙ぐぶ",
			charset: charset.RGBY.Japanese,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if d := tC.charset.Decode(tC.encoded); d != tC.decoded {
				t.Fail()
			}
		})
	}
}

func TestSkipsUnknownCharacters(t *testing.T) {
	e := []byte{1, 146, 135, 80}
	cs := charset.RGBY.English
	d := cs.Decode(e)

	if d != "SH" {
		t.Fail()
	}
}
