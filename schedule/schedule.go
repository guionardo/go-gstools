package schedule

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Schedule represents a scheduled event
// You can use the builder pattern to create a schedule:
// schedule := NewSchedule("test").Every(300 * time.Millisecond)
type Schedule struct {
	Name          string
	Data          map[string]any
	runsEvery     time.Duration
	runsAfter     time.Time
	enabled       bool
	LastRun       time.Time
	dontRunBefore time.Duration
	dontRunAfter  time.Duration
	runCount      uint64
}

// Creates a new basic schedule
func NewSchedule(name string) *Schedule {
	return &Schedule{
		Name:      name,
		Data:      make(map[string]any),
		runsEvery: time.Duration(0),
		runsAfter: time.Time{},
		enabled:   true,
		LastRun:   time.Time{},
	}
}

// Setups schedule to run every duration
func (s *Schedule) Every(duration time.Duration) *Schedule {
	s.runsEvery = duration
	return s
}

// Setups schedule to run after time
func (s *Schedule) After(t time.Time) *Schedule {
	s.runsAfter = t
	return s
}

// Setups schedule to not run before duration
// Duration represents a time of day without date. 
// You can use schedule.ParseDuration to create a duration
func (s *Schedule) DontRunBefore(duration time.Duration) *Schedule {
	s.dontRunBefore = duration
	return s
}

// Setups schedule to not run after duration
// Duration represents a time of day without date. 
// You can use schedule.ParseDuration to create a duration
func (s *Schedule) DontRunAfter(duration time.Duration) *Schedule {
	s.dontRunAfter = duration
	return s
}

// Returns the time of the next run or empty if not enabled or no next run
func (s *Schedule) NextRun() time.Time {
	if !s.enabled {
		return time.Time{}
	}
	if s.runsEvery > 0 {
		return s.getNextRunWindow(s.LastRun.Add(s.runsEvery))
	}
	if s.runsAfter.After(time.Now()) {
		return s.getNextRunWindow(s.runsAfter)
	}
	return time.Time{}
}

func (s *Schedule) Enabled(enabled bool) *Schedule {
	s.enabled = enabled
	return s
}

func (s *Schedule) getNextRunWindow(nextRun time.Time) time.Time {
	if s.dontRunBefore > 0 {
		dontRunBefore := Today().Add(s.dontRunBefore)
		if time.Now().Before(dontRunBefore) {
			return Tomorrow().Add(s.dontRunBefore)
		}
	}

	if s.dontRunAfter > 0 {
		dontRunAfter := Today().Add(s.dontRunAfter)
		if time.Now().After(dontRunAfter) {
			return Tomorrow().Add(s.dontRunAfter)
		}
	}

	return nextRun
}

func (s *Schedule) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Schedule('%s'", s.Name))
	if s.runsEvery > 0 {
		sb.WriteString(fmt.Sprintf(", runs every %v", s.runsEvery))
	}
	if !s.runsAfter.IsZero() {
		sb.WriteString(fmt.Sprintf(", runs after %v", s.runsAfter))
	}
	if !s.enabled {
		sb.WriteString(", disabled")
	}
	if !s.LastRun.IsZero() {
		sb.WriteString(fmt.Sprintf(", last run %v", s.LastRun))
	}
	if !s.NextRun().IsZero() {
		sb.WriteString(fmt.Sprintf(", next run %v", s.NextRun()))
	}
	sb.WriteString(")")
	return sb.String()
}

func (s *Schedule) RunCount() uint64 {
	return s.runCount
}

func (s *Schedule) Run(callback ScheduleCallBack, logger *log.Logger) {
	callback(s)
	s.LastRun = time.Now()
	s.runCount++
	if logger != nil {
		logger.Printf("Running schedule: %v", s)
	}
}
