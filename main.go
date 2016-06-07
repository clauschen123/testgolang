package main

import (
	"fmt"
)

func main() {
	//	fmt.Println("called by go")
	//	var str string
	//	fmt.Scan(&str)
	test_base()
}

//---------------------- test struct ----------------------
//public struct
type Student struct {
	Name  string //`学生姓名`
	Age   int    //`a:"1111 "b:"3333"` //这个不是单引号，而是~键上的符号
	class string
}

func (this Student) getName() string      { return this.Name }
func (this *Student) setName(name string) { this.Name = name }
func (this *Student) getAge() int         { return this.Age }
func (this *Student) setAge(age int)      { this.Age = age }
func (this *Student) Display()            { fmt.Println(this.Name, ",", this.Age) }
func (this *Student) GetName() string     { return this.Name }
func (this *Student) GetAge() int         { return this.Age }
func (this *Student) PrintName()          { fmt.Println(this.Name) }
func (this *Student) GetAge1() int        { return this.Age }

//匿名字段
type CollegeStudent struct {
	Student
	Profession string
}

//接口是一系列操作的集合，是一种约定。我们可以把它看作与其它对象通讯的协议。 任
//何非接口类型只要拥有某接口的全部方法，就表示它实现了该接口，
type IStudent interface {
	GetName() string
	GetAge() int
}
