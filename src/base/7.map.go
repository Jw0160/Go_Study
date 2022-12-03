package main

import "fmt"

/**
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
    map[KeyType]ValueType
		KeyType:表示键的类型。
		ValueType:表示键对应的值的类型。
*/
func main() {
	//BaseDemo()
	//InitDemo()
	//foreachDemo(map[string]int{
	//	"1": 1,
	//	"2": 2,
	//})
	delDemo()

}

/**
通过make创建 map
*/
func BaseDemo() {
	m := make(map[string]int, 2)
	m["test1"] = 90
	m["test2"] = 100
	fmt.Println(m)
}

/**
通过 map 直接构造数据
*/
func InitDemo() {
	m := map[string]int{}
	m["hello2"] = 200
	fmt.Println(m)
	key, b := hasKey(m, "hello2")
	fmt.Println(key, b)
}

/**
如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
*/
func hasKey(m map[string]int, k string) (int, bool) {
	value, ok := m[k]
	return value, ok

}

func foreachDemo(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}

}

func delDemo() {
	m := map[string]int{
		"1": 1,
		"2": 2,
	}
	fmt.Println(m)
	delete(m, "3")
	fmt.Println(m)
}
