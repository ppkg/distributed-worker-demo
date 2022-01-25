package handler

import (
	"fmt"
	"time"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/kit"
)

type subPlugin struct {
}

func (s *subPlugin) Name() string {
	return "sub"
}

func (s *subPlugin) Handle(Id int64, jobId int64, input string) (string, error) {
	var params multiParam
	_ = kit.JsonDecode([]byte(input), &params)
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("%s->完成任务(%d,%d)\n", s.Name(), Id, jobId)
	return kit.JsonEncode(map[string]interface{}{
		"result": params.Result - 1,
	}), nil
}

func NewSub() core.PluginHandler {
	return &subPlugin{}
}
