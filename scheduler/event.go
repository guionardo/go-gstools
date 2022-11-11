package scheduler

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	Action         func(context.Context) error
	ScheduledEvent struct {
		id                string
		runAfter          time.Time
		runEvery          time.Duration
		runCount          int
		lastRun           time.Time
		deactivateOnError bool
		enabled           bool
		action            Action
		lock              sync.RWMutex
		parent            *Scheduler
		lastError         error
		afterRun          func(ctx context.Context, event *ScheduledEvent) error
	}
)

func NewEvent(action Action, parent *Scheduler, options ...ScheduledEventOption) *ScheduledEvent {
	event := &ScheduledEvent{
		id:       uuid.New().String(),
		action:   action,
		parent:   parent,
		enabled:  true,
		runAfter: time.Now(),
	}
	event.setOptions(options...)
	return event
}

func (e *ScheduledEvent) RunCount() int {
	return e.runCount
}

func (e *ScheduledEvent) Disable() {
	e.enabled = false
}

func (e *ScheduledEvent) updateNextRun() {
	var logLine string

	if !e.enabled {
		logLine = "disabled"
	} else if e.runAfter.IsZero() {
		e.runAfter = time.Now()
		logLine = "first run"
	} else if e.runEvery > 0 {
		e.runAfter = e.runAfter.Add(e.runEvery)
		logLine = fmt.Sprintf("next run in %v", e.runAfter)
	} else {
		e.enabled = false
		logLine = "disabled after one run"
	}
	e.parent.Logf("%s %s", e, logLine)
}

func (e *ScheduledEvent) String() string {
	return fmt.Sprintf("#%s @ %v", e.id, e.NextRun())
}

func (e *ScheduledEvent) NextRun() (nextRun time.Time) {
	if !e.enabled {
		return
	}

	if e.runEvery > 0 {
		nextRun = e.lastRun.Add(e.runEvery)
	} else {
		nextRun = e.runAfter
	}

	return
}

func (e *ScheduledEvent) setId(id string) *ScheduledEvent {
	if id != "" && e.id != id {
		e.id = id
	}
	return e
}

func (s *ScheduledEvent) GetRunner() func(context.Context) error {
	return func(ctx context.Context) error {
		if !s.enabled {
			return nil
		}
		nextRun := s.NextRun()
		if nextRun.After(time.Now()) {
			s.parent.Logf("%s waiting for next run in %v", s, nextRun.Sub(time.Now()))
			time.Sleep(nextRun.Sub(time.Now()))
		}
		s.Run(ctx)
		return nil
	}
}

func (s *ScheduledEvent) Run(ctx context.Context, runAfter ...func()) {
	go func() {
		s.lock.Lock()
		defer s.lock.Unlock()
		logs := make([]string, 0, 4)
		defer func() {
			for _, fn := range runAfter {
				fn()
			}
			s.parent.Logf("%s", strings.Join(logs, ", "))
			if s.afterRun != nil {
				s.afterRun(ctx, s)
			}
		}()

		if s.runEvery > 0 && time.Now().Sub(s.lastRun) < s.runEvery {
			logs = append(logs, "skip")
			return
		}

		logs = append(logs, fmt.Sprintf("RUN %s", s))

		if !s.enabled {
			logs = append(logs, "disabled")
			return
		}

		s.runCount++
		logs = append(logs, fmt.Sprintf("runCount %d - last %v", s.runCount, time.Now().Sub(s.lastRun)))
		s.lastRun = time.Now()
		if s.lastError = s.action(ctx); s.lastError != nil {
			logs = append(logs, fmt.Sprintf("error: %v", s.lastError))
			if s.deactivateOnError {
				s.enabled = false
			}
		}
		s.updateNextRun()
	}()
}
