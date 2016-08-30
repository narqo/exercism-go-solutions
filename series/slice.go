package slice

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var l []string
	if n > len(s) {
		return l
	}
	for i := 0; i < len(s) && i+n <= len(s); i++ {
		l = append(l, s[i:i+n])
	}
	return l
}

func First(n int, s string) (string, bool) {
	l := All(n, s)
	if len(l) == 0 {
		return "", false
	}
	return l[0], true
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	f, _ := First(n, s)
	return f
}
