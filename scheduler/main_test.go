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
func TestScheduler_Run(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		testDuration := 10 * time.Second
		log.Printf("Starting test - duration = %v", testDuration)

		s := NewScheduler().SetLogger(log.Default())
		var n1, n2, n3 int
		e1 := s.AddEvent(func(ctx context.Context) error {
			n1++
			return nil
		}, RunEvery(1*time.Second), Id("e1"))
		c1 := int(testDuration.Seconds() / e1.runEvery.Seconds())
		e2 := s.AddEvent(func(ctx context.Context) error {
			n2++
			return nil
		}, RunAfter(time.Now().Add(500*time.Millisecond)), Id("e2"))
		c2 := 1 // e2 should run only once
		e3 := s.AddEvent(func(ctx context.Context) error {
			n3++
			return nil
		}, RunEvery(2*time.Second), Id("e3"))
		c3 := int(testDuration.Seconds() / e3.runEvery.Seconds())

		ctx, cancel := context.WithTimeout(context.Background(), testDuration)
		defer cancel()
		s.Run(ctx)

		if e1.RunCount() != c1 {
			t.Errorf("n1 should be %d, but got %d", c1, n1)
		}
		if e2.RunCount() != c2 {
			t.Errorf("n2 should be %d, but got %d", c2, n2)
		}
		if e3.RunCount() != c3 {
			t.Errorf("n3 should be %d, but got %d", c3, n3)
		}
	})

}

func TestScheduler_RunUntilEnd(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		testDuration := 5 * time.Second
		log.Printf("Starting test - duration = %v", testDuration)

		s := NewScheduler().SetLogger(log.Default())
		var n1 int
		e1 := s.AddEvent(func(ctx context.Context) error {
			n1++
			return nil
		},
			RunEvery(100*time.Millisecond),
			Id("e1"),
			AfterRunFunction(func(ctx context.Context, event *ScheduledEvent) error {
				log.Printf("AfterRunFunction called for event %v", event)
				if event.RunCount() > 3 {
					event.enabled = false
				}
				return nil
			}))
		c1 := 4

		s.RunUntilEmpty()

		if e1.RunCount() != c1 {
			t.Errorf("n1 should be %d, but got %d", c1, n1)
		}

	})

}
