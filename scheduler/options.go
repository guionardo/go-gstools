package scheduler

import (
	"context"
	"fmt"
	"time"
)

type ScheduledEventOption struct {
	Type  int
	Value interface{}
}

const (
	tParent = iota
	tId
	tRunAfter
	tEvery
	tDeactivateOnError	
	tAfterRunFunction
)

func Parent(parent *Scheduler) ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tParent,
		Value: parent,
	}
}

func Id(id string) ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tId,
		Value: id,
	}
}

func RunAfter(t time.Time) ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tRunAfter,
		Value: t,
	}
}

func RunEvery(t time.Duration) ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tEvery,
		Value: t,
	}
}

func DeactivateOnError() ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tDeactivateOnError,
		Value: true,
	}
}

func AfterRunFunction(f func(ctx context.Context, event *ScheduledEvent) error) ScheduledEventOption {
	return ScheduledEventOption{
		Type:  tAfterRunFunction,
		Value: f,
	}
}

func (event *ScheduledEvent) setOptions(options ...ScheduledEventOption) {
	var logs = make([]string, len(options))
	for i, option := range options {
		switch option.Type {
		case tParent:
			event.parent = option.Value.(*Scheduler)
			logs[i] = fmt.Sprintf("parent=%v", event.parent)

		case tId:
			event.setId(option.Value.(string))
			logs[i] = fmt.Sprintf("id=%v", event.id)

		case tRunAfter:
			event.runAfter = option.Value.(time.Time)
			logs[i] = fmt.Sprintf("after=%v", event.runAfter)

		case tEvery:
			event.runEvery = option.Value.(time.Duration)
			logs[i] = fmt.Sprintf("interval=%v", event.runEvery)

		case tDeactivateOnError:
			event.deactivateOnError = option.Value.(bool)
			logs[i] = fmt.Sprintf("deactivateOnError=%v", event.deactivateOnError)

		case tAfterRunFunction:
			event.afterRun = option.Value.(func(ctx context.Context, event *ScheduledEvent) error)
			logs[i] = "AfterRunFunction"
		}
	}
	if event.runEvery == 0 {
		logs = append(logs, "runOnce")
	}
	event.parent.Logf("%s options = %v", event, logs)
}
