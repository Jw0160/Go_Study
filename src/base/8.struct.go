package main

import (
	"encoding/json"
	"fmt"
)

/**

结构体中字段大写开头表示可公开访问，小写表示私有
    type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
其中：
    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
    3.字段类型：表示结构体字段的具体类型。

type person struct {
        name string
        city string
        age  int8
    }
*/
func main() {
	//person := Person{
	//	name: "Tome",
	//	city: "changsha",
	//	age:  20,
	//}
	//var person1 Person
	//person1.city = "zhuzhou"
	//
	//fmt.Println(person)
	//fmt.Println(person1)
	//
	//FooStruct()
	//PointStruct()
	//extendsDemo()
	PrintJsonDemo()

}

/**
类别名
*/
//类型定义
type NewInt int

//类型别名
type MyInt = int

func AliasDemo() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

}

//结构体
type Person struct {
	name, city string
	age        int
}

//匿名结构体
func FooStruct() {
	var foo = struct {
		name string
	}{"fooStruct11"}
	fmt.Println(foo)

	var foo1 struct {
		name1 string
	}
	foo1.name1 = "foo1"

	fmt.Println(foo1)

}

//需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
func PointStruct() {
	i := new(int)
	// i = 1 // is error
	*i = 1
	//创建指针类型结构体
	p := new(Person)
	p.age = 10
	fmt.Println(p, &p, *p)
	fmt.Println(i, &i, *i)
	//取结构体的地址实例化
	p3 := &Person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
}

//结构体的匿名字段  --- 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
type PersonNoName struct {
	string
	int
}

func NiMingStruct() {
	p1 := PersonNoName{
		"pprof.cn",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"pprof.cn", int:18}
	fmt.Println(p1.string, p1.int) //pprof.cn 18
}

// 嵌套匿名结构体 (继承?)

//嵌套结构体的字段名冲突,为了避免歧义需要指定具体的内嵌结构体的字段。 不指定则使用子类的字段
type Animal struct {
	name string
}

func (receiver *Animal) say() {
	fmt.Printf("Animal name is %v\n", receiver.name)
}

type Dog struct {
	Animal
	age  int
	name string
}

func extendsDemo() {
	dog := Dog{
		age: 0,
	}
	dog.name = "xiaogou"
	dog.Animal.name = "father animal"
	fmt.Printf("%#v\n", dog)

	dog.say()

}

//结构体标签（Tag）
/**
Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。

Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
 `key1:"value1" key2:"value2"`

结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
键值对之间使用一个空格分隔。
注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
例如不要在key和value之间添加空格。
*/
type JsonDemo struct {
	ID      int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name    string //json序列化是默认使用字段名作为key
	address string //私有不能被json包访问
}

func PrintJsonDemo() {
	jsonDemo := JsonDemo{
		ID:      1,
		Name:    "tom",
		address: "herr",
	}
	marshal, _ := json.Marshal(jsonDemo)
	fmt.Printf("%s\n", marshal) //{"id":1,"Name":"tom"}
	jsonDemo.address = "12312312321"
	fmt.Printf("%s\n", jsonDemo.address) //{"id":1,"Name":"tom"}
}

type HeroModel struct {
	AtkGrow   float64  `json:"atk_grow"`
	Command   int      `json:"command"`
	ArmTypes  []string `json:"arm_types"`
	Atk       int      `json:"atk"`
	BaseSkill struct {
		FitArms []int `json:"fit_arms"`
		UseHero struct {
		} `json:"use_hero"`
		DeDesc   string `json:"de_desc"`
		WorkType int    `json:"work_type"`
		Person   []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"person"`
		Chance      int    `json:"chance"`
		Name        string `json:"name"`
		InheritHero struct {
		} `json:"inherit_hero"`
		Quality   int `json:"quality"`
		SkillType int `json:"skill_type"`
		Id        int `json:"id"`
	} `json:"base_skill"`
	Magic        int     `json:"magic"`
	Id           int     `json:"id"`
	Speed        int     `json:"speed"`
	Quality      int     `json:"quality"`
	Group        int     `json:"group"`
	Level        int     `json:"level"`
	CommandGrow  float64 `json:"command_grow"`
	GetWay       []int   `json:"get_way"`
	Politics     int     `json:"politics"`
	Control      int     `json:"control"`
	CharmGrow    float64 `json:"charm_grow"`
	PoliticsGrow float64 `json:"politics_grow"`
	Filename     string  `json:"filename"`
	Charm        int     `json:"charm"`
	SpeedGrow    float64 `json:"speed_grow"`
	MagicGrow    float64 `json:"magic_grow"`
	Season       int     `json:"season"`
	Name         string  `json:"name"`
	Targets      []int   `json:"targets"`
	InheritSkill struct {
		FitArms []int `json:"fit_arms"`
		UseHero struct {
		} `json:"use_hero"`
		DeDesc   string `json:"de_desc"`
		WorkType int    `json:"work_type"`
		Person   struct {
		} `json:"person"`
		Chance      int    `json:"chance"`
		Name        string `json:"name"`
		InheritHero []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"inherit_hero"`
		Quality   int `json:"quality"`
		SkillType int `json:"skill_type"`
		Id        int `json:"id"`
	} `json:"inherit_skill"`
}

//
//type student struct {
//	name string
//	age  int
//}
//
//func main() {
//	m := make(map[string]*student)
//	stus := []student{
//		{name: "pprof.cn", age: 18},
//		{name: "博客", age: 28},
//		{name: "测试", age: 23},
//	}
//	fmt.Println(stus)
//
//	m[stus[0].name] = &stus[0]
//	m[stus[1].name] = &stus[1]
//	m[stus[2].name] = &stus[2]
//
//以下方式为何会输出失败数据 ?
//	//for _, stu := range stus {
//	//
//	//	a := &stu
//	//	fmt.Printf("a = %v\n", a)
//	//	fmt.Printf("stu.name = %v\n", stu.name)
//	//	m[stu.name] = a
//	//}
//
//	fmt.Println(m["博客"])
//	fmt.Println(m["测试"])
//	fmt.Println(m["pprof.cn"])
//
//	for k, v := range m {
//		fmt.Println(k, "=>", v.name)
//	}
//}
