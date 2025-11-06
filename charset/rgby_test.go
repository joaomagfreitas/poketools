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
			desc:    "decodes using english charset",
			encoded: []byte{128, 146, 135, 80, 137, 128, 130, 138, 80, 141, 132},
			decoded: "ASH",
			charset: charset.RGBY.English,
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
