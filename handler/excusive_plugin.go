package handler

import (
	"fmt"
	"time"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/kit"
)

type exclusivePlugin struct {
}

func (s *exclusivePlugin) Name() string {
	return "exclusive"
}

func (s *exclusivePlugin) Handle(Id int64, jobId int64, input string) (string, error) {
	var params []multiParam
	_ = kit.JsonDecode([]byte(input), &params)
	var result int
	for _, item := range params {
		result += item.Result
	}
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("%s->完成任务(%d,%d)\n", s.Name(), Id, jobId)
	return kit.JsonEncode(map[string]interface{}{
		"result": result + 5,
	}), nil
}

func NewExclusive() core.PluginHandler {
	return &exclusivePlugin{}
}
