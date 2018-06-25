// 定义包名 main表示本文件为可独立执行程序,每一个Go应用程序都应该包含一个名为main的包
package main

import "fmt" //引入包 fmt 包实现了格式化 IO（输入/输出）的函数

func main() { //func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
	fmt.Println("Hello World!")               // 打印函数后跟\n
	fmt.Print("Hello World!\n")               // 打印函数
	fmt.Printf("%d %s!", 2018, "Hello World") // 格式化输出
}

// 执行 go run hello.go
