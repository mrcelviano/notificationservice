package scheduler

import (
	"github.com/robfig/cron/v3"
)

const cronSpec = "0 * * * * *"

func Start(runFunc func()) error {
	cr := cron.New(cron.WithSeconds())
	_, err := cr.AddFunc(cronSpec, runFunc)
	if err != nil {
		return err
	}

	cr.Start()
	return nil
}
