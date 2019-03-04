// Package gigasecond calculates the time after 10e9 seconds
package gigasecond

import "time"

const (
	gigasecond time.Duration = 1000000000
)

// AddGigasecond returns the time after 10e9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigasecond)
}
