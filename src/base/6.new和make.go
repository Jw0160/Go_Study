package main

import "fmt"

func main() {
	//FirstTest()
	//newDemo2()
	//makeDemo()
	Test()
}

// FirstTest panic: runtime error: invalid memory address or nil pointer dereference
//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
///*
func FirstTest() {
	var a *int
	//i = new(int)  使用new 为 指针分配内存后,才能赋值
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	//b = make(map[string]int,10) 对于 map slice channel 使用make 分配内存后,才能使用
	b["测试"] = 100
	fmt.Println(b)
}

/**
  func new(Type) *Type
使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
*/
func newDemo() {
	i := new(int8)
	fmt.Printf("%t %T\n", i, i) //%!t(*int8=0xc0000aa058) *int8
}

func newDemo2() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Printf("%v %v\n", i, *i)
}

/**
	func make(t Type, size ...IntegerType) Type
只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
*/
func makeDemo() {
	var b map[string]int
	b = make(map[string]int, 5)
	b["测试"] = 100
	b["测试1"] = 100
	b["测试2"] = 100
	b["测试3"] = 100
	b["测试4"] = 100
	b["测试5"] = 100
	fmt.Println(b)
}

/**
程序定义一个int变量num的地址并打印
将num的地址赋给指针ptr，并通过ptr去修改num的值
*/
func Test() {
	var num int
	fmt.Printf("%v \n", &num)
	ptr := &num
	*ptr = 100
	fmt.Printf("%v \n", num)
}
