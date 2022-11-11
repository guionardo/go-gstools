package scheduler

import (
	"context"
	"log"
	"time"
)

type Scheduler struct {
	collection *EventCollection
	logger     *log.Logger
	running    bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		collection: NewEventCollection(),
	}
}

func (s *Scheduler) AddEvent(action Action, options ...ScheduledEventOption) *ScheduledEvent {
	event := NewEvent(action, s, options...)
	s.collection.Add(event)

	return event
}

func (s *Scheduler) Run(ctx context.Context) {
	s.Logf("Starting scheduler")
	defer s.Logf("Stopped scheduler")
	s.running = true

	for {
		if s.collection.IsEmpty() {
			s.Logf("No more events to run")
			return
		}
		next, waitTime := s.collection.GetNext()

		select {
		case <-ctx.Done():
			return

		case <-time.After(waitTime):
			if next != nil {
				next.Run(ctx)
			}
		}
	}
}

func (s *Scheduler) RunUntilEmpty() {
	s.Run(context.Background())
}
