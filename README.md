# golangdemo
## 搭建Golang环境可能遇到的问题
### 需要把git目录C:\Program Files\Git\bin添加到系统变量path中
go: missing Git command. See https://golang.org/s/gogetcmd
package github.com/gin-gonic/gin: exec: "git": executable file not found in %PATH%
### CMD下执行下面命令即可
go get github.com/go-sql-driver/mysql
go get github.com/astaxie/beego
go get github.com/gin-gonic/gin

### goland运行项目
runnerw.exe: CreateProcess failed with error 216 (no message available)
修改包名为 package main 就行。 不然执行main方法报错

