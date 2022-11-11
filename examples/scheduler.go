package examples

import (
	"context"
	"log"
	"time"

	"github.com/guionardo/go-gstools/scheduler"
)

func scheduler_example() {
	sch := scheduler.NewScheduler()
	sch.AddEvent(func(ctx context.Context) error {
		println("Hello")
		return nil
	},
		scheduler.Id("my-event"),
		scheduler.RunEvery(1*time.Second),
		scheduler.AfterRunFunction(func(ctx context.Context, event *scheduler.ScheduledEvent) error {
			log.Printf("Runnin event %s - #%d", event, event.RunCount())
			if event.RunCount() > 3 {
				event.Disable()
			}
			return nil
		}))
	sch.RunUntilEmpty()
}
