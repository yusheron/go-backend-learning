# Day01 - Go 语言基础

## 1. Go 程序入口

Go 程序从 `main` 包中的 `main` 函数开始执行。
只有包含 `func main()` 的程序才能被编译为可执行文件。

最小可执行示例：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

## 2. 变量声明

Go 使用 `var` 或 `:=` 来声明变量，类型在编译期确定。

三种常见声明方式：

```go
var a int = 10      // 显式声明类型
var b = 20          // 类型推导
c := 30              // 简短声明
```

## 3. 基本数据类型
Go 是强类型语言，变量在声明时就确定了数据类型。

常用的基本数据类型包括：

- `int`：整数类型
- `float64`：浮点数类型
- `bool`：布尔类型
- `string`：字符串类型

示例：

```go
var age int = 20
var score float64 = 95.5
var passed bool = true
var name string = "Go"
```

## 4. if 条件判断
Go 的 `if` 语句用于条件判断，不需要使用括号包裹条件。

基本用法：

```go
age := 20
if age >= 18 {
    fmt.Println("成年人")
}
```

## 5. for 循环
Go 语言只有一种循环结构：`for`，可以表示多种循环形式。

最常见的形式：

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

## 6. 函数定义
Go 使用 `func` 关键字定义函数，参数类型写在参数名之后。

基本函数定义示例：

```go
func add(a int, b int) int {
    return a + b
}
```

