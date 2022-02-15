package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ppkg/distributed-worker/core"
	"github.com/ppkg/distributed-worker/dto"
	"github.com/ppkg/kit"
)

func TestAsyncSubmit(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "asyncSubmit",
		},
	}
	appCtx := core.NewApp(core.WithNacosAddrOption("mse-e52dbdd6-p.nacos-ans.mse.aliyuncs.com:8848"), core.WithNacosNamespaceOption("27fdefc2-ae39-41fd-bac4-9256acbf97bc"), core.WithNacosServiceGroupOption("my-service"))
	appCtx.Init()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rpcReq := dto.AsyncJobRequest{
				Name: "async-job异步job2",
				Meta: map[string]string{
					"author": "zihua123",
					"qq":     "597291437",
				},
				Label: "sn0001",
				PluginSet: []string{
					"plus",
					"multi",
				},
				Type:     "test",
				IsNotify: true,
			}
			for i := 0; i < 100; i++ {
				rpcReq.TaskInputList = append(rpcReq.TaskInputList, kit.JsonEncode(map[string]interface{}{
					"a": 1,
					"b": 2,
				}))
			}
			startTime := time.Now()
			rw := os.Stderr
			fmt.Fprintf(rw, "开始发送异步请求，时间:%s，数据量：%d \n", startTime.Format("2006-01-02 15:04:05"), len(rpcReq.TaskInputList))
			jobId, err := appCtx.AsyncSubmitJob(rpcReq)
			endTime := time.Now()
			if err != nil {
				fmt.Fprintf(rw, "-------------------------------------------------------------\n")
				fmt.Fprintf(rw, "异步请求异常退出，时间:%s，耗时：%f秒，err:%+v \n", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime).Seconds(), err)
				return
			}

			fmt.Fprintf(rw, "-------------------------------------------------------------\n")
			fmt.Fprintf(rw, "异步请求完成，时间:%s，耗时：%f秒，jobId:%d\n", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime).Seconds(), jobId)
		})
	}
}
