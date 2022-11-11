package scheduler

import (
	"log"
)

func (s *Scheduler) SetLogger(logger *log.Logger) *Scheduler {
	s.logger = logger
	return s
}

func (s *Scheduler) Logf(format string, v ...interface{}) {
	if s.logger != nil {
		s.logger.Printf(format, v...)
	}
}
