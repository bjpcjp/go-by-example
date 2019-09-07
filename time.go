/* time.go */

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// current time

	now := time.Now()
	p(now)

	// building a time structure

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// monday-sunday workday support
	p(then.Weekday())

	// comparisons
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// returns a duration (interval)
	diff := now.Sub(then)
	p(diff)

	// durations in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// advancing (or regressing) a time by a duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
