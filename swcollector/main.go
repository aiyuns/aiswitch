package main

import (
	"flag"
	"fmt"
	"os"

	"aiyun.com.cn/aiswitch/swcollector/cron"
	"aiyun.com.cn/aiswitch/swcollector/funcs"
	"aiyun.com.cn/aiswitch/swcollector/g"
	"aiyun.com.cn/aiswitch/swcollector/http"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	funcs.BuildMappers()

	cron.ReportAgentStatus()
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()
	cron.Collect()

	go http.Start()

	select {}

}