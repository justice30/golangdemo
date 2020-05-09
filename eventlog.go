package main

import (
	"fmt"
	_ "github.com/elastic/beats/libbeat/common"
	_ "github.com/elastic/beats/winlogbeat/eventlog"
	//"golang.org/x/sys/windows/svc/eventlog"
	"os"
)

func main() {
	fmt.Println("爱媛媛")
	s := os.Environ()
	fmt.Println(s)
	//var option common.Config
	//var elog, err = eventlog.New((*common.Config)(option))
	//if err == nil{
	//	fmt.Println("eventlog异常")
	//	return
	//}
	//var name = elog.Name();
	//fmt.Println(name)
	//var elog, err = eventlog.Open("jinlong")
	//if err != nil {
	//	return
	//}
	//defer elog.Close()
	//print()
	//elog.Info(1, "金龙测试:爱金金！ 应用程序日志： 信息")
	//elog.Warning(2, "金龙测试:爱金金！ 应用程序日志： 警告")
	//elog.Error(3, "金龙测试:爱金金！ 应用程序日志： 错误")
	//println(err)

}
