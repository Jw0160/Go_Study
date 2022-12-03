package main

import "fmt"

/**
标准声明
  var 变量名 变量类型

*/
func normalVar() {
	var s1 string
	fmt.Println(s1)
}

/**
批量声明
    var (
        a string
        b int
        c bool
        d float32
    )
*/

func batchVar() {
	var (
		s1 string
		i1 int
		b1 bool
	)
	fmt.Println(s1, b1, i1)

}

/*
byte和rune类型
  uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。

    rune类型，代表一个 UTF-8字符。
*/

func byteAndRune() {
	// 遍历字符串
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

/**
变量的初始化
    var 变量名 类型 = 表达式
    var name, sex = "pprof.cn", 1

:= 方式声明并初始化变量
    n := 10
    m := 200 // 此处声明局部变量m
*/
func main() {
	normalVar()
	batchVar()
	byteAndRune()
	//多行字符串使用反引号 `
	s1 := `
dasd
asdaslj
adas`

	fmt.Println(s1)
}
