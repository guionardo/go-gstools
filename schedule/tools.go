package schedule

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Get the current time of day without date
func WallClock() time.Duration {
	now := time.Now()
	return now.Sub(Today())
}

// Get the date of today without time
func Today() time.Time {	
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// Get the date of tomorrow without time
func Tomorrow() time.Time {
	return Today().Add(24 * time.Hour)
}

func validateNumber(str string, min, max int) (n int, err error) {
	n, err = strconv.Atoi(str)
	if err == nil && (n < min || n > max) {
		err = fmt.Errorf("invalid number %d, must be between %d and %d", n, min, max)
	}
	return
}

// Parse a duration string in the format "hh:mm:ss" or "hh:mm" or in duration string (see time.ParseDuration)
func ParseDuration(str string) (duration time.Duration, err error) {
	duration, err = time.ParseDuration(str)
	if err == nil {
		return
	}

	w := strings.Split(str, ":")
	if len(w) < 2 {
		err = errors.New("invalid duration")
		return
	}
	var h, m, s int
	if h, err = validateNumber(w[0], 0, 23); err != nil {
		return
	}
	if m, err = validateNumber(w[1], 0, 59); err != nil {
		return
	}
	if len(w) > 2 {
		if s, err = validateNumber(w[2], 0, 59); err != nil {
			return
		}
	}

	duration = time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
	return
}
