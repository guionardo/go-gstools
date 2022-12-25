package schedule

import (
	"context"
	"log"
	"sync"
	"time"
)

type (
	Scheduler struct {
		schedules   []*Schedule
		runInterval time.Duration
		logger      *log.Logger
		lock        sync.RWMutex
	}
	ScheduleCallBack func(*Schedule)
)

func NewScheduler() *Scheduler {
	return &Scheduler{
		schedules:   []*Schedule{},
		runInterval: time.Second,
	}
}

func (s *Scheduler) SetInterval(interval time.Duration) *Scheduler {
	s.runInterval = interval
	return s
}

func (s *Scheduler) SetLogger(logger *log.Logger) *Scheduler {
	s.logger = logger
	return s
}

func (s *Scheduler) Log(format string, v ...interface{}) {
	if s.logger != nil {
		s.logger.Printf(format, v...)
	}
}

func (s *Scheduler) AddSchedule(schedule *Schedule) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Log("Adding schedule: %v", schedule)
	s.schedules = append(s.schedules, schedule)
}

func (s *Scheduler) GetNextSchedule() *Schedule {
	s.lock.Lock()
	defer s.lock.Unlock()

	var next *Schedule
	for _, schedule := range s.schedules {
		if !schedule.enabled {
			continue
		}
		if next == nil {
			next = schedule
			continue
		}
		nextRun := schedule.NextRun()
		if !nextRun.IsZero() && nextRun.Before(next.NextRun()) {
			next = schedule
		}
	}
	if next != nil {
		s.Log("Next schedule: %v", next)
	}
	return next
}

func (s *Scheduler) RunWithChannel(ctx context.Context, channel chan *Schedule) {
	s.Run(ctx, func(schedule *Schedule) {
		channel <- schedule
	})
}

func (s *Scheduler) Run(ctx context.Context, callback ScheduleCallBack) {

	for {
		select {
		case <-ctx.Done():
			s.Log("Scheduler context done")
			return
		default:
			next := s.GetNextSchedule()

			if next != nil && next.NextRun().Before(time.Now()) {
				next.Run(callback, s.logger)
			} else {
				time.Sleep(s.runInterval)
			}
		}
	}
}
