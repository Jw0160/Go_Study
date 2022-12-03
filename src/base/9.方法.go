package main

/**
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
       函数体
   }
    1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
    3.方法名、参数列表、返回参数：具体格式与函数定义相同。


指针类型的接收者由一个结构体的指针组成，
由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
这种方式就十分接近于其他语言中面向对象中的this或者self。

当方法作用于值类型接收者时，
Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

 什么时候应该使用指针类型接收者
    1.需要修改接收者中的值
    2.接收者是拷贝代价比较大的大对象
    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。


在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
    fmt.Println("Hello, 我是一个int。")
}
func main() {
    var m1 MyInt
    m1.SayHello() //Hello, 我是一个int。
    m1 = 100
    fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
*/
func main() {

}

type PersonFuc struct {
	name string
	age  int8
}

func (p *PersonFuc) Name() string {
	return p.name
}

func (p *PersonFuc) SetName(name string) {
	p.name = name
}

func (p *PersonFuc) Age() int8 {
	return p.age
}

func (p *PersonFuc) SetAge(age int8) {
	p.age = age
}
