package secret

func Handshake(code int) []string {
	if code < 0 {
		return nil
	}
	var h []string
	if code&1 != 0 {
		h = append(h, "wink")
	}
	if code&2 != 0 {
		h = append(h, "double blink")
	}
	if code&4 != 0 {
		h = append(h, "close your eyes")
	}
	if code&8 != 0 {
		h = append(h, "jump")
	}

	if code&16 != 0 {
		for left, right := 0, len(h)-1; left < right; left, right = left+1, right-1 {
			h[left], h[right] = h[right], h[left]
		}
	}

	return h
}
