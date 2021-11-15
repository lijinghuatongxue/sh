package main

import "fmt"

type Human struct {
	name  string
	age   string
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}
type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

// SayHi Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Printf("la la la la... %s", lyrics)
}

// SayHi Employee重写Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Men Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	JingHua := Student{Human{"Jinghua", "18", "Iphone"}, "JiWang", 267}
	liuqi := Employee{Human{"LiuQi", "15", "Iphone"}, "QianXun", 30000}
	//定义Men类型的变量i
	var i Men

	i = JingHua
	fmt.Println("This is Jinghua, a Student:")
	i.SayHi()
	i.Sing("追梦赤子心")
	i = liuqi
	fmt.Println("This is liuqi, an Employee:")
	i.SayHi()
	i.Sing("有为歌")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 2)
	//T这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1] = JingHua, liuqi
	for _, value := range x {
		value.Sing("hei hei hei..")
	}

}

//$ go run human_student_employee_demo.go
//This is Jinghua, a Student:
//Hi, I am Jinghua you can call me on Iphone
//la la la la... 追梦赤子心This is liuqi, an Employee:
//Hi, I am LiuQi, I work at QianXun. Call me on Iphone
//la la la la... 有为歌Let's use a slice of Men and see what happens
//la la la la... hei hei hei..la la la la... hei hei hei..%
