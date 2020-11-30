# Google 资深工程师深度讲解 Go 语言

## 切片

+ slice 可以向后扩展，但不能向前扩展

+ 取下标由 len 决定: `s[i], i < len(s)`

+ reslice 由 cap 决定：`s[n:m], m>=n, m < cap(s)`