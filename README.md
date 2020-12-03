# Google 资深工程师深度讲解 Go 语言

> macOS 11.0.1 / Debian 10, Go 1.15.5, LiteIDE 37.3

## 切片

### 扩展

+ slice 可以向后扩展，但不能向前扩展

+ 取下标由 len 决定: `s[i], i < len(s)`

+ reslice 由 cap 决定：`s[n:m], m>=n, m < cap(s)`

### 添加新元素

+ 当添加元素时超出 cap, 系统会分配新的更大的底层数组

+ 由于值传递的关系，必须接收 `append` 的返回值

## map

### map 的 Key

+ 使用哈希表，所以 Key 必须是可比较的

+ 除了 slice, map, function 之外的内置类型都可以作map的Key

+ 如果 struct 不包含上述类型的字段，也可以作为 key

## 面向“对象”

+ go 语言只支持封装，不支持继承和多态

+ 没有 class, 只有 struct

+ nil 指针也可以调用方法

## 面向接口

+ 接口变量自带指针

+ 接口变量同样采用值传递，几乎不需要使用接口的指针

+ 指针接收都实现只能以指针方式使用；值接收者都可以

## 函数式编程

### 函数式编程 VS 函数指针

+ 函数是一等公民：参数、变量、返回值都可以是函数

+ 高阶函数

+ 闭包

### "正统" 函数式编程

+ 不可变性：不能有状态，只有常量和函数

+ 函数只能有一个参数

## 错误处理和资源管理

### defer 调用

+ 确保调用在函数结束时发生

+ 参数在defer语句时计算 

+ defer列表为后进先出

### panic

+ 停止当前函数执行

+ 一直向上返回，执行每一层的 defer

+ 如果没有遇见 recover，程序退出

### recover

+ 仅在 defer 调用中使用

+ 获取 panic 的值 

+ 如果无法处理，可重新 panic

### error vs panic

+ 意料之中的：使用error。如：文件打不开

+ 意料之外的：使用painc。如：数组越界

## 测试与性能调优

### 传统测试与表格驱动测试

#### 传统测试

+ 测试数据和测试逻辑混在一起

+ 出错信息不明确 

+ 一旦一个数据出错，测试全部结束

#### 表格驱动测试

+ 分离的测试数据与测试逻辑

+ 明确的出错信息

+ 可以部分失败

+ go语言的语法使得我们更易实践表格驱动测试

## goroutine

### 协程 coroutine

+ 轻量级“线程”

+ ***非抢占式***多任务处理，由协程主动交出控制权

+ 编译器/解释器/虚拟机层面的多任务

+ 多个协程可能在一个或多个线程上运行

### goroutine 的定义

+ 任何函数只需要加上`go`就能送给调度器运行

+ 不需要在定义时区分是否是异步函数

+ 调试器在合适的点进行切换

+ 使用 `-rece` 来检测数据访问冲突。如：`go run -rece main.go`

### goroutine 可能的切换点

+ I/O, select

+ channel

+ 等待锁

+ 函数调用(有时)

+ `runtime.Gosched()`