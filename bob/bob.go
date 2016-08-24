package bob

import "strings"

const testVersion = 2

// Hey returns Bob's answer to msg.
func Hey(msg string) string {
	msg = strings.TrimSpace(msg)
	if len(msg) == 0 {
		return "Fine. Be that way!"
	}
	if strings.ToUpper(msg) == msg && strings.ToLower(msg) != msg {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(msg, "?") {
		return "Sure."
	}
	return "Whatever."
}
