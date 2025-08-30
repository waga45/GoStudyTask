package task2

import (
	"fmt"
	"math"
)

type Shape interface {
	//面积
	Area() float64
	//周长
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Rd float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Rd * c.Rd
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Rd
}

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
func CalculateAreaAndRect() {
	var r = Rectangle{Width: 10, Height: 10}
	area := r.Area()
	perimeter := r.Perimeter()
	fmt.Println("Rectangle面积：", area, "，周长：", perimeter)

	var c = Circle{Rd: 10}
	area1 := c.Area()
	perimeter1 := c.Perimeter()
	fmt.Println("Circle面积：", area1, "，周长：", perimeter1)
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("员工ID：", e.EmployeeID, "，姓名：", e.Name, "，年龄：", e.Age)
}

/*
*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
func ComposeObject() {
	var e = Employee{EmployeeID: 100000, Person: Person{Name: "张三", Age: 20}}
	e.PrintInfo()
}
