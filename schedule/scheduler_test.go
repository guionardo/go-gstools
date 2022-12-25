package schedule

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewScheduler_(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		scheduler := NewScheduler().SetInterval(100 * time.Millisecond).SetLogger(logger)

		s1 := NewSchedule("test").Every(300 * time.Millisecond)
		s2 := NewSchedule("test2").Every(500 * time.Millisecond)
		scheduler.AddSchedule(s1)
		scheduler.AddSchedule(s2)

		ns := scheduler.GetNextSchedule()
		if ns.Name != "test" {
			t.Errorf("GetNextSchedule() = %v, want %v", ns.Name, "test")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		scheduler.Run(ctx, func(schedule *Schedule) {
			t.Logf("Running schedule: %v", schedule)
		})

		if s1.RunCount() != 4 {
			t.Errorf("s1.Runs() = %v, want %v", s1.RunCount(), 4)
		}
		if s2.RunCount() != 2 {
			t.Errorf("s2.Runs() = %v, want %v", s2.RunCount(), 2)
		}

	})

}


