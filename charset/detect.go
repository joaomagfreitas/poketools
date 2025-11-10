package charset

// Detect resolves the most optimal charset for decoding a
// a set of data blocks.
// It sets [ok] flag false, if none of the charset can decode the data blocks.
func Detect(data [][]byte, cs ...Charset) (ok bool, charset Charset) {
	if len(data) == 0 {
		return
	}

	if len(cs) == 0 {
		return
	}

	t := make([]int, len(cs))
	for c := range cs {
		for _, d := range data {
			if !cs[c].Decodable(d) {
				continue
			}

			t[c]++
		}
	}

	csi := 0
	h := t[csi]
	for i, t := range t {
		if t > h {
			h = t
			csi = i
		}
	}

	if h == 0 {
		return
	}

	ok = true
	charset = cs[csi]
	return
}
