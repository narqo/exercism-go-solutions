package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

const gigasecond = 1e9 * time.Second

// AddGigasecond add 1 Gs to the `in` time.
func AddGigasecond(in time.Time) time.Time {
	return in.Add(gigasecond)
}
