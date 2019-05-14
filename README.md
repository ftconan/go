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













