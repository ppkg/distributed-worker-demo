package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ppkg/distributed-worker/dto"
	"github.com/ppkg/kit"
)

func regAsyncRequest() {
	http.HandleFunc("/async", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("content-type", "text/html; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		rpcReq := dto.SyncJobRequest{
			Name: "async-job异步job",
			PluginSet: []string{
				"plus",
				"multi",
			},
		}
		for i := 0; i < 5000; i++ {
			rpcReq.TaskInputList = append(rpcReq.TaskInputList, kit.JsonEncode(map[string]interface{}{
				"a": 1,
				"b": 2,
			}))
		}
		startTime := time.Now()
		fmt.Fprintf(rw, "开始发送异步请求，时间:%s，数据量：%d <br/>", startTime.Format("2006-01-02 15:04:05"), len(rpcReq.TaskInputList))
		resp, err := app.SyncSubmitJob(rpcReq)
		endTime := time.Now()
		if err != nil {
			fmt.Fprintf(rw, "-------------------------------------------------------------<br>")
			fmt.Fprintf(rw, "异步请求异常退出，时间:%s，耗时：%f秒，err:%+v <br/>", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime).Seconds(), err)
			return
		}

		fmt.Fprintf(rw, "-------------------------------------------------------------<br>")
		fmt.Fprintf(rw, "异步请求完成，时间:%s，耗时：%f秒，job状态:%d，返回结果:%s <br/>", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime).Seconds(), resp.Status, resp.Result)

		fmt.Fprintf(rw, "-------------------------------------------------------------<br>")
		fmt.Fprintf(rw, "最终计算结果：%d <br/>", computeResult(resp.Result))
	})
}
