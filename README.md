# go-learning
### 第一章(ch1)
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
### 第二章(ch2)
* 变量赋值(与其他主要语言差异)
    1. 赋值可以进行自动类型推断
    2. 在一个赋值语句中可以对多个变量进行同时赋值
* 常量定义(与其他主要语言差异)
    1. 快速设置连续值(iota)
* 基本数据类型
    1. bool
    2. string
    3. int int8 int16 int32 int64
    4. uint uint8 uint16 uint32 uint64 uintptr
    5. byte    // alias for uint8
    6. rune    // alias for int32, represents a Unicode code point
    7. float32 float64
    8. complex64 complex128
### 第三章(ch3)
* 类型转换(与其他主要语言差异)
    1. Go语言不允许隐式类型转换
    2. 别名和原有类型也不能进行隐式类型转换
* 类型预定义值
    1. math.MaxInt64
    2. math.MaxFloat64
    3. math.MaxUint32
* 指针类型(与其他主要语言差异)
    1. 不支持指针运算
    2. string是值类型,其默认初始值为空字符串,而不是nil
### 第四章(ch4)
* 算数运算符
    1. \+ - * / % ++ --
    2. Go语言没有前置++, --
* 比较运算符
    1. == != > < >= <=
* 用==比较数组
    1. 相同维数且含有相同个数元素的数组才可以比较
    2. 每个元素都相同的才相等
* 逻辑运算符
    1. && || !
* 位运算符(与其他主要语言差异)
    1. & | ^ << >>
    2. &^ 按位置0
### 第五章(ch5)
* 循环(与其他主要语言差异)
    1. Go语言仅支持循环关键字
    2. for j := 7; j <= 9; j++
* if条件(与其他主要语言差异)
    1. condition表达式结果必须为布尔值
    2. 支持变量赋值:
    3. if var declaration; condition {
        
    }
* switch(与其他主要语言差异)
    1. 条件表达式不限制为常量或者整数
    2. 单个case中,可以出现多个结果选项,使用逗号分隔;
    3. 与C语言规则相反,Go不需要break来明确退出一个case;
    4. 可以不设定switch之后的条件表达式,在此种情况下,整个switch结构与多个if...else...的逻辑等同
### 第六章(ch6)
* 数组声明
    1. var a [3]int          //声明并初始化为0
    2. b := [3]int{1, 2, 3}   // 声明同时初始化
    3. c := [2][2]int{{1, 2}, {3, 4}} // 多维数组初始化
    4. 相同维数相同长度可以比较
* 数组截取(与其他主要语言差异)
    1. a[开始索引(包含), 结束索引(不包含)]
    2. 不支持负数索引
* 切片(可变长数组)
    1. 指针,len,cap
    2. 不可比较,只能和nil比较
### 第七章(ch7)
* map
    1. s := map[string]int{"a": 1, "b": 2}
### 第八章(ch8)
* map(extend)
    1. map的value可以是一个方法
    2. 与Go的Dock type接口方式一起,可以方便实现单一方法对象的工厂模式
* 实现Set(与其他主要语言差异)
    1. Go的内置集合中没有Set实现,可以map[type]bool
        * 元素唯一性
        * 基本操作
            1. 添加元素
            2. 判断元素是否存在
            3. 删除元素 
            4. 元素个数
### 第9章(ch9)
* 字符串(与其他主要语言差异)
    1. string是数据类型,不是引用和指针类型
    2. string是只读的byte slice, len函数是它所包含的byte数
    3. string的byte数组可以存放任何数据
    4. string是不可变的byte slice
* Unicode UTF8
    1. Unicode是一种字符集(code point)
    2. UTF8是unicode的存储实现(转换为字节序列的规则)
* 常用的字符串函数
    1. strings http://golang.org/pkg/strings/
    2. strconv http://golang.org/pkg/strconv/
### 第十章(ch10)
* 函数是一等公民(与其他主要语言差异)
    1. 可以有多个返回值
    2. 所有参数都是值传递：slice,map,channel会有传引用的错觉
    3. 函数可以作为变量的值
    4. 函数可以作为参数和返回值    
    5. 函数可变参数相当于数组
    6. defer函数
### 第十一章(ch11)
* 行为的定义和实现
    1. 定义
    2. 行为
* 接口(与其他主要语言差异)
    1. 接口为非入侵性的,实现不依赖于接口定义
    2. 接口的定义可以包含在接口使用者包内
    3. 自定义类型
### 第十二章(ch12)
* 扩展和复用
### 第十三章(ch13)
* 多态
* 空接口和断言
    1. 空接口可以表示任何类型
    2. 通过断言将空接口转换为制定类型
    3. v, ok := p.(int) // ok=true 转换成功
* Go接口最佳实践
    1. 倾向于使用小的接口定义,很多接口只包含一个方法
    2. 较大的接口定义,可以有多个小接口组合而成
    3. 只依赖于必要功能的最小接口
### 第十四章(ch14)
* Go的错误机制(与其他主要语言差异)
    1. 没有异常处理
    2. error类型实现了error接口
    3. 可以通过errors.New来快速创建错误实例
* Go错误最佳实践
    1. 定义不同错误变量,以便于判断错误类型
    2. 及早失败,避免嵌套
* panic
    1. panic用于不可恢复的错误
    2. panic退出前会执行defer指定的内容
* panic vs os.Exit
    1. os.Exit退出时不会调用defer指定的函数
    2. os.Exit退出时不输出当前调用栈信息
* recover
    1. 形成僵尸服务进程,导致health check失效
    2. "Let's it Crash!"往往是我们恢复不确定性错误的最好办法
### 第十五章(ch15)
* 构建可复用模块
    1. 基本复用模块单元(以首字母大写来表明可被包外代码访问)
    2. 代码的package可以和所在的目录不一致
    3. 同一目录里的Go代码的package要保持一致  
* init方法
    1. 在main被执行前,所有依赖的package的init方法都会被执行
    2. 不同包的init函数按照包导入的依赖关系决定执行顺序
    3. 每个包可以有多个init函数
    4. 包的每个源文件也可以有多个init函数,这点比较特殊
* package
    1. 通过go get来获取远程依赖
    2. go get -u 强制从远程更新依赖(go get -u github.com/easierway/concurrent_map)
    3. 注意代码在github上的组织形式,以适应go get
    4. 直接以代码路径开始,不要有src
* Go未解决问题
    1. 同一环境,不同项目使用同一包的不同版本
    2. 无法管理对包的特定版本的依赖
    3. vendor(查找包路径解决方案)
        1. 当前包下的vendor目录
        2. 向上级目录查找,直到找到src下的vendor目录
        3. 在GOPATH下面查找依赖包
        4. 在GOROOT目录下查找
    4. 常用的依赖管理工具
        1. godep http://github.com/tools/godep
        2. glide http://github.com/Masterminds/glide
        3. dep http://github.com/golang/dep 
### 第十六章(ch16)
* Thead vs Groutine
    1. 创建时默认的stack的大小
        1. JDK5以后Java Thread stack 默认为1M
        2. Groutine 的Stack初始化大小2k
    2. 和KSE(Kernel Space Entity)的对应关系
        1. Java Thread 1：1
        2. Groutine    M : N 
  




