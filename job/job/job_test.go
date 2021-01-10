package job

import (
	"github.com/CH3CHO/GoMultiModuleTest/core/service"
	"github.com/golang/mock/gomock"
	"github.com/golobby/container/pkg/container"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCreateJobs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	shService := NewMockSayHelloService(mockCtrl)
	shService.EXPECT().SayHello(gomock.Any()).DoAndReturn(func(name string) string {
		return "[Mocked]Hello " + name + "!"
	})

	ctn := container.NewContainer()
	ctn.Singleton(func() service.SayHelloService {
		return shService
	})

	jobs := CreateJobs(ctn)

	assert.Equal(t, 1, len(jobs))

	job := jobs[0]
	assert.Equal(t, "*job.sayHelloJob", reflect.TypeOf(job).String())
	job.Run()
}
