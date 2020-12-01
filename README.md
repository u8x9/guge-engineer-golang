# Google 资深工程师深度讲解 Go 语言

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