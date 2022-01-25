package handler

import (
	"fmt"
	"time"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/kit"
)

type plus2Plugin struct {
}

func (s *plus2Plugin) Name() string {
	return "plus2"
}

func (s *plus2Plugin) Handle(Id int64, jobId int64, input string) (string, error) {
	var params multiParam
	_ = kit.JsonDecode([]byte(input), &params)
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("%s->完成任务(%d,%d)\n", s.Name(), Id, jobId)
	return kit.JsonEncode(map[string]interface{}{
		"result": params.Result + 5,
	}), nil
}

func NewPlus2() core.PluginHandler {
	return &plus2Plugin{}
}
