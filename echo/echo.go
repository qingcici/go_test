package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// func main() {
// 	//切片
// 	//echo1 显式
// 	// var s, sep string
// 	// for i := 1; i < len(os.Args); i++ { //符号:=是短变量声明，i++是语法，不像C是表达式
// 	// 	s += sep + os.Args[i] //连接字符串sep和os.Args,运算符+=是赋值运算符
// 	// 	sep = " "
// 	// }
// 	// fmt.Println(s)

// 	//echo2
// 	// s, sep := "", ""
// 	// for _, arg := range os.Args[1:] {
// 	// 	s += sep + arg
// 	// 	sep = " "
// 	// }
// 	// fmt.Println(s)

// 	// echo3
// 	// fmt.Println(strings.Join(os.Args[1:], " "))
// 	// fmt.Println(os.Args[1:])

// 	// 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
// 	fmt.Println(strings.Join(os.Args[:], " "))

// 	//修改echo程序，使其打印每个参数的索引和值，每个一行。
// 	for index, arg := range os.Args[1:] {
// 		fmt.Println(index+1, ":", arg)
// 	}
// }

//做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异
//go run echo.go 1 3
func main() {
	joinEcho()
	plusEcho()
}

func joinEcho() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("echo2: %fs\n", time.Since(start).Seconds())
}

func plusEcho() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("echo1: %fs\n", time.Since(start).Seconds())
}
