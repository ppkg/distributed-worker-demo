package handler

import (
	"fmt"
	"time"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/kit"
)

type multiPlugin struct {
}

func (s *multiPlugin) Name() string {
	return "multi"
}

func (s *multiPlugin) Handle(Id int64, jobId int64, input string) (string, error) {
	var params multiParam
	_ = kit.JsonDecode([]byte(input), &params)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%s->完成任务(%d,%d)\n", s.Name(), Id, jobId)
	return kit.JsonEncode(map[string]interface{}{
		"result": params.Result * 10,
	}), nil
}

func NewMulti() core.PluginHandler {
	return &multiPlugin{}
}

type multiParam struct {
	Result int
}
