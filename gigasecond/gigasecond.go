package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	// reviewers may protest!
	return t.Add(time.Second * time.Duration(math.Pow(10, 9)))
}
