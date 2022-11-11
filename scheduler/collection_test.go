package scheduler

import (
	"context"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Ltime)
}
func TestNewEventCollection(t *testing.T) {
	f := func(ctx context.Context) error {
		return nil
	}

	t.Run("Default", func(t *testing.T) {
		c := NewEventCollection()
		s := NewScheduler().SetLogger(log.Default())

		e1 := NewEvent(f, s, RunEvery(1*time.Second), Id("e1"))
		e2 := NewEvent(f, s, RunEvery(500*time.Millisecond), Id("e2"))
		e3 := NewEvent(f, s, RunAfter(time.Now().Add(3*time.Second)), Id("e3"))

		c.Add(e1, e2, e3)

		next, runAfter := c.GetNext()
		if next != e2 {
			t.Errorf("next should be %v, but got %v", e2, next)
		}
		if runAfter > 500*time.Millisecond {
			t.Errorf("runAfter should be less than 500ms, but got %v", runAfter)
		}
		c.Run(e2.id, context.Background())

		next, _ = c.GetNext()
		if next != e1 {
			t.Errorf("after e2 run, next should be %v, but got %v", e1, next)
		}

	})

}
