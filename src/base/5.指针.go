package main

import "fmt"

/**
区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
3个概念：指针地址、指针类型和指针取值。
2个符号：&（取地址）和*（根据地址取值）


ptr := &v    // v的类型为T
	v:代表被取地址的变量，类型为T
	ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
*/
func main() {
	//PointDefault()
	PointEmpty()
}

func PointDefault() {
	i := 10
	pi := &i

	fmt.Printf("%t\n", i)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%t\n", *pi)
}

/**
Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
*/
func PointEmpty() {

	var p *string //当一个指针被定义后没有分配到任何变量时，它的值为 nil
	fmt.Printf("%t\n", p)
	*p = "100" //panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("%t\n", p)

}
