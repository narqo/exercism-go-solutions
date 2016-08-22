package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

const gigasecond time.Duration = 1000000000 * time.Second

// AddGigasecond calculates the date that someone turned or will celebrate their 1 Gs anniversary.
func AddGigasecond(in time.Time) time.Time {
	return in.Add(gigasecond)
}
