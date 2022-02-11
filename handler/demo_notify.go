package handler

import (
	"encoding/json"
	"fmt"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/distributed-worker/dto"
	"github.com/ppkg/distributed-worker/enum"
	"github.com/ppkg/kit"
)

type demoNotify struct {
}

// 返回handler名称，与job中type字段对应
func (s *demoNotify) Name() string {
	return "test"
}

// 通知业务处理
func (s *demoNotify) Handle(data dto.JobNotify) error {
	if data.Status != enum.FinishJobStatus {
		fmt.Printf("job(%d)执行失败,失败原因:%s \n", data.Id, data.Message)
		return nil
	}
	var list []dataItem
	_ = json.Unmarshal([]byte(data.Result), &list)
	var result int
	for _, item := range list {
		result += item.Result
	}
	fmt.Printf("job(%d)执行成功,计算结果:%d \n", data.Id, result)
	return nil
}

// job开始执行通知
func (s *demoNotify) PostStart(data dto.StartNotify) error {
	fmt.Println("任务开始执行了.....", kit.JsonEncode(data))
	return nil
}

func NewDemoNotify() core.JobNotifyHandler {
	return &demoNotify{}
}

type dataItem struct {
	Result int
}
