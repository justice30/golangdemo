package main
import (
	"fmt"
	"golang.org/x/sys/windows/svc/eventlog"
)

func main() {
	fmt.Println("爱金金")
	var elog, err = eventlog.Open("jinlong")
	if err != nil {
		return
	}
	defer elog.Close()
	print()
	elog.Info(1, "金龙测试:爱金金！ 应用程序日志： 信息")
	elog.Warning(2, "金龙测试:爱金金！ 应用程序日志： 警告")
	elog.Error(3, "金龙测试:爱金金！ 应用程序日志： 错误")
	println(err)

}