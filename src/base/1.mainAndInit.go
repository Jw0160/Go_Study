package main

import "fmt"

func init() {
	fmt.Println("init 0")
}
func init() {
	fmt.Println("init 2")
}
func init() {
	fmt.Println("init 1")
}

/**
对同一个go文件的init()调用顺序是从上到下的。
对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。
对于不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，最后调用main函数。
如果init函数中使用了println()或者print()你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用。
*/

/**
	相同点：
        两个函数在定义时"不能有任何的参数和返回值"，且Go程序自动调用。
    不同点：
        init可以应用于任意包中，且可以重复定义多个。
        main函数只能用于main包中，且只能定义一个。
*/
func main() {
	fmt.Println("main")
}
