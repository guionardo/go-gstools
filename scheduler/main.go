package scheduler

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	Scheduler struct {
		notify chan int
		events map[string]*ScheduledEvent
	}
	ScheduledEvent struct {
		id                string
		after             time.Time
		runOnce           bool
		lastRun           time.Time
		deactivateOnError bool
		enabled           bool
		action            func(context.Context) error
	}
)

func (e *ScheduledEvent) RunOnce() *ScheduledEvent {
	e.runOnce = true
	return e
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		notify: make(chan int),
		events: make(map[string]*ScheduledEvent, 10),
	}
}

func (s *Scheduler) getNextEvent() *ScheduledEvent {
	var next *ScheduledEvent
	for _, event := range s.events {
		if next == nil || event.after.Before(next.after) {
			next = event
		}
	}
	return nil
}

func (s *Scheduler) Execute(ctx context.Context, event *ScheduledEvent) {
	//TODO: Implementar executação da action
	//
	event.action(ctx)
}

func (s *Scheduler) AddEventAfterTime(execTime time.Time, action func(context.Context) error) {
	id := uuid.New().String()
	event := &ScheduledEvent{
		id:     id,
		after:  execTime,
		action: action,
	}

	s.events[event.id] = event
}

func (s *Scheduler) Run(ctx context.Context) error {
	for {
		next := s.getNextEvent()
		var waitTime time.Duration

		if next == nil {
			waitTime = time.Second

		} else {
			waitTime = next.after.Sub(time.Now())
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(waitTime):
			if next != nil {
				s.Execute(ctx, next)

			}
		case <-s.notify:
			if event := s.getNextEvent(); event != nil {
				if err := event.action(ctx); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
