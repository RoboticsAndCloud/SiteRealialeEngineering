/* main.go - the main function */
/*
modification history
--------------------
Fei
*/
/*
DESCRIPTION
*/
package main

import (
	"flag"
	"fmt"
	"os"
)



import (
//	"mini_spider"
)

var (
//	spider   *mini_spider.Spider
	logPath  *string = flag.String("l", "./log", "日志路径")
	confPath *string = flag.String("c", "./conf", "配置路径")
	version  *bool   = flag.Bool("v", false, "版本信息")
)

const QUEUE_SIZE = 10000

type versionConfig struct {
	version string
}

/**
 * Description: 启动Spider
 */

func f()(string){
	return "test"
}

func main(){
//	config.LoadVersion(*confPath)

	flag.Parse()
	fmt.Printf("mini_spider: version %s\n", f())
	fmt.Printf("logPath:%v \n", *logPath);
	fmt.Printf("confPath:%v \n", *confPath);

	if *version {
		fmt.Printf("mini_spider: version 1.0\n")
		Exit(0)
	}
	str := "abc"

	point := &str
	fmt.Println(*point)

}


func Exit(errno int) {
	os.Exit(errno)
}

/*
func main() {
	var (
		errno int = 0
		err   error
	)
	flag.Parse()
	if *version {
		fmt.Printf("mini_spider: version 1.0\n")
		Exit(0)
	}
	if err = log.Init("mini-spider", "DEBUG", *logPath, true, "midnight", 7); err != nil {
		errno = -1
		goto ERR
	}
	if spider, err = mini_spider.NewSpider(*confPath, QUEUE_SIZE); err != nil {
		log.Logger.Error("mini_spider.NewSpider fails")
		errno = -1
		goto ERR
	}

	//启动spider
	spider.Run()
	//等待spider处理完所有任务
	spider.Wait()

ERR:
	log.Logger.Close()
	Exit(errno)
}
*/