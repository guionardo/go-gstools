## Agendador (mais simples)

Executar tarefas agendadas ou periódicas em segundo plano

## Setup

```bash
❯ go get github.com/guionardo/go-gstools
go: downloading github.com/guionardo/go-gstools v0.9.4
go: added github.com/guionardo/go-gstools v0.9.4
```

On your project, you have to setup the schedules.

```golang
package main

import (
	"context"
	"fmt"
    "log"
	"os"
	"os/signal"
	"time"

	"github.com/guionardo/go-schedule/schedule"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	scheduler := schedule.NewScheduler().SetLogger(logger)
	// Add a task that runs every 15 minutes
	scheduler.AddSchedule(schedule.NewSchedule("First Task").Every(15 * time.Minute))
	// Add a task that runs every 20 minutes between 06:00 and 08:00
	scheduler.AddSchedule(schedule.NewSchedule("Second Task").Every(20 * time.Minute).
		DontRunBefore(6 * time.Hour).
		DontRunAfter(8 * time.Hour))

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	scheduler.Run(ctx, func(sch *schedule.Schedule) {
		fmt.Printf("Running %v", sch)
	})

}
```

The first example shows the event loop blocked, but you can start as a goroutine and control the exit with your context.

You, also, can use the ```scheduler.RunWithChannel(ctx, channel)``` to populate your channel and control the loop event in your application.