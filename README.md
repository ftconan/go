# go-learning
* 开发环境 GOPATH
    1. go 1.8 以后Unix默认$HOME/go
    2. Windows默认%USERPROFILE%/go
    3. Mac上修改～/.bash_profile设置
* go command
    1. go version               (查看go版本)
    2. go run hello_world.go    (运行go)
    3. go build hello_world.go  (编译go) 
* 应用程序入口
    1. 必须是main包: package main
    2. 必须是main方法： func main()
    3. 文件名不一定是main.go
* 退出返回值(与其他语言区别)
    1. Go中main函数不支持任何返回值
    2. 通过os.Exit返回状态
* 获取命令行参数(与其他语言区别)
    1. main函数不支持传入参数func main(arg []string)
    2. 在程序中直接通过os.Args获取命令行参数
* 编写测试程序
    1. 源码文件以_test结尾：xxx_test.go
    2. 测试方法名以Test开头:func TestXXX(t *testing.T){}
* 变量赋值(与其他主要语言差异)
    1.赋值可以进行自动类型推断
    2.在一个赋值语句中可以对多个变量进行同时赋值
* 常量定义(与其他主要语言差异)
    1.快速设置连续值(iota)
