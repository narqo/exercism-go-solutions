package bob

import "strings"

const testVersion = 2 // same as targetTestVersion

func Hey(msg string) string {
	str := strings.TrimSpace(msg)
	if len(str) == 0 {
		return "Fine. Be that way!"
	}
	if !onlyNums(str) && strings.ToUpper(str) == str {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(str, "?") {
		return "Sure."
	}
	return "Whatever."
}

func onlyNums(str string) bool {
	normalized := strings.Map(func(r rune) rune {
		if r < 65 {
			return -1
		}
		return r
	}, str)
	return len(normalized) == 0
}
