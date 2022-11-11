package scheduler

import (
	"context"
	"sync"
	"time"
)

type EventCollection struct {
	events  map[string]*ScheduledEvent
	running sync.Map
	lock    sync.RWMutex
}

func NewEventCollection() *EventCollection {
	return &EventCollection{
		events: make(map[string]*ScheduledEvent, 10),
	}
}

func (c *EventCollection) Add(events ...*ScheduledEvent) {
	for _, event := range events {
		c.events[event.id] = event
	}
}

func (c *EventCollection) Remove(event *ScheduledEvent) {
	delete(c.events, event.id)
}

func (c *EventCollection) IsEmpty() bool {
	for _, event := range c.events {
		if event.enabled {
			return false
		}
	}
	return true
}

func (c *EventCollection) GetNext() (event *ScheduledEvent, waitTime time.Duration) {
	var next *ScheduledEvent

	for _, event := range c.events {
		if c.IsRunning(event.id) || !event.enabled {
			continue
		}
		if next == nil || event.NextRun().Before(next.NextRun()) {
			next = event
		}
	}
	if next == nil {
		waitTime = time.Second
		event = nil
	} else {
		event = next
		runAfter := next.NextRun()
		if runAfter.Before(time.Now()) {
			waitTime = time.Millisecond
		} else {
			waitTime = runAfter.Sub(time.Now())
		}
	}
	return
}

func (c *EventCollection) IsRunning(eventId string) bool {
	if _, ok := c.running.Load(eventId); ok {
		return true
	}
	return false
}

func (c *EventCollection) Run(eventId string, ctx context.Context) {
	if event, ok := c.events[eventId]; ok {
		if c.IsRunning(eventId) {
			return
		}
		c.running.Store(eventId, 0)
		go func() {
			event.Run(ctx, func() { c.running.Delete(eventId) })
		}()
	}

}
