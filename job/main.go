package main

import (
	"github.com/CH3CHO/GoMultiModuleTest/core/service"
	"github.com/CH3CHO/GoMultiModuleTest/job/job"
	"github.com/golobby/container/pkg/container"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	ctn := container.NewContainer()
	service.InitContainer(&ctn)

	jobs := job.CreateJobs(ctn)

	for _, j := range jobs {
		_, _ = c.AddFunc(j.Spec(), func() {
			j.Run()
		})
	}

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
