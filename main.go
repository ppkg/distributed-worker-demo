package main

import (
	"distributed-worker-demo/handler"
	"flag"
	"fmt"
	"net/http"

	"github.com/ppkg/distributed-worker/core"
)

var (
	port         = flag.Int("port", 8001, "TCP port for this node")
	isEnableHttp = flag.Bool("http", false, "is enable http server")
	app          *core.ApplicationContext
)

func main() {
	flag.Parse()

	if *isEnableHttp {
		go http.ListenAndServe(":8080", nil)
	}

	app = core.NewApp(core.WithNacosAddrOption("mse-e52dbdd6-p.nacos-ans.mse.aliyuncs.com:8848"), core.WithNacosNamespaceOption("27fdefc2-ae39-41fd-bac4-9256acbf97bc"), core.WithNacosServiceGroupOption("my-service"), core.WithPortOption(*port))

	// 注册task处理插件
	app.RegisterPlugin(func(ctx *core.ApplicationContext) core.PluginHandler {
		return handler.NewPlus()
	}).RegisterPlugin(func(ctx *core.ApplicationContext) core.PluginHandler {
		return handler.NewMulti()
	})

	// 注册回调通知
	app.RegisterJobNotify(func(ctx *core.ApplicationContext) core.JobNotifyHandler {
		return handler.NewDemoNotify()
	})

	err := app.Run()
	if err != nil {
		fmt.Println("start app got err:", err)
	}
}
