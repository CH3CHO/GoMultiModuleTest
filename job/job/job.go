package job

import (
	"fmt"
	"github.com/CH3CHO/GoMultiModuleTest/core/service"
	"github.com/golobby/container/pkg/container"
	"time"
)

type Job interface {
	Spec() string
	Run()
}

type sayHelloJob struct {
	service service.SayHelloService
}

func (s *sayHelloJob) Spec() string {
	return "@every 1s"
}

func (s *sayHelloJob) Run() {
	now := time.Now()
	name := now.Format(time.RFC3339)
	text := s.service.SayHello(name)
	fmt.Println(text)
}

func CreateJobs(ctn container.Container) []Job {
	var jobs []Job

	var shService service.SayHelloService
	ctn.Make(&shService)
	shJob := sayHelloJob{service: shService}
	jobs = append(jobs, &shJob)

	return jobs
}
