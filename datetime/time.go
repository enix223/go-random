package datetime

import (
	"math/rand"
	"time"
)

// RandomDate get a random date between from and to
func RandomDate(from, to time.Time) time.Time {
	if to.Before(from) {
		panic("to date must after from date")
	}

	d := to.Sub(from)
	g := rand.Int63n(d.Nanoseconds())
	return from.Add(time.Duration(g))
}
