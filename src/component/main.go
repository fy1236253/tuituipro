package main

import (
	"cfg"
	"component/cron"
	"component/g"
	"component/http"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"redis"
	"syscall"

	"github.com/astaxie/beego/logs"
)

func main() {
	cfgString := flag.String("c", "cfg.json", "specify config file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	cfg.ParseConfig(*cfgString) //配置文件
	g.InitWxConfig()            //微信相关参数
	// g.InitDB()           //db池
	g.InitRootDir()      //全局参数
	redis.InitConnPool() //redis 链接初始化

	//logTo := cfg.Config().Logs
	//if logTo != "stdout" {
	//	f, err := os.OpenFile(logTo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//	if err != nil {
	//		panic(fmt.Sprintf("open logfile error"))
	//	}
	//	defer f.Close()
	//	log.SetOutput(f)
	//}
	loggers := logs.NewLogger(1024)
	loggers.SetLogger("tuike", `{"filename":"logs/log.log"}`)
	// 日志追加pid和时间
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.SetPrefix(fmt.Sprintf("PID.%d ", os.Getpid()))
	go http.Start()
	go cron.Start()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("all service stopping...")
		// cron.Stop()
		// http.Stop()
		// proc.Stop()

		// mq.ConnPool.Close() // 关闭连接池
		redis.ConnPool.Close()
		log.Println("all service stop ok ")
		os.Exit(0)
	}()
	select {}
}
