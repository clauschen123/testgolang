package main

import "fmt"

//Test of slice
func test_slice() {
	//声明一个2个元素的数组，名字为 arr_1,因为是 int 型数组，所以初值为0，即[0,0]
	var arr_1 [2]int
	/*声明一个2个元素的数组，名字为 arr_2，并同时赋初值，{}里为空，说明没有赋初
	  值，等同于上面*/
	arr_2 := [2]int{}
	//声明一个2个元素的数组，名字为 arr3, arr3_1, arr3_2，并同时赋初值，结果均为[1,2]
	arr3 := [2]int{1, 2}
	//{}里的冒号左边是下标，右边是值
	arr3_1 := [2]int{0: 1, 1: 2}
	arr3_2 := [2]int{1: 2, 0: 1}
	/*不指定数组长度，自动计算长度, [...],声明一个2个（自动计算而来）元素的数组，名字
	  为 arr4，并同时赋初值，结果为[1,2]*/
	arr4 := [...]int{1, 2}
	/*声明一个4个（自动计算而来）元素的数组，名字为 shuzu5，并同时赋初值，结果
	  为[0,0,0,9]*/
	arr5 := [...]int{3: 9}
	fmt.Println(arr_1)
	fmt.Println(arr_2)
	fmt.Println(arr3)
	fmt.Println(arr3_1)
	fmt.Println(arr3_2)
	fmt.Println(arr4)
	fmt.Println(arr5)
}

//Test of struct,inherit,interface
type Student struct {
	Name  string
	Age   int
	class string
}

func (this Student) getName() string {
	return this.Name
}
func (this *Student) setName(name string) {
	this.Name = name
}
func (this *Student) getAge() int {
	return this.Age
}
func (this *Student) setAge(age int) {
	this.Age = age
}
func (this *Student) Display() {
	fmt.Println(this.Name, ",", this.Age)
}

type CollegeStudent struct {
	Student
	Profession string
}

type IStudent interface {
	GetName() string
	GetAge() int
}

func (this *Student) GetName() string {
	return this.Name
}
func (this *Student) GetAge() int {
	return this.Age
}

func test_struct() {

	var si IStudent = &Student{"李四abcv", 23, "2004(2)班"}
	fmt.Println(si.GetName())
	return

	s1 := new(Student)
	s1.Name = "张三"
	s1.Age = 12
	s1.class = "21班"
	fmt.Println(s1)

	s2 := Student{"张三", 12, "21 班"}
	fmt.Println(s2)

	s3 := Student{Name: "张三", Age: 12, class: "22 班"}
	fmt.Println(s3)

	s := Student{Name: "张三", Age: 15, class: "32班"}
	fmt.Println(s.getName(), s.getAge())

	s.setName("claus")
	s.setAge(100)
	fmt.Println(s.getName(), s.getAge())

	s4 := CollegeStudent{Student: Student{Name: "李四", Age: 23, class: "2004(2)班"},
		Profession: "物理"}
	s4.Display()
	fmt.Println(s4.Student.Name) //可以通过 student 访问 Name
	fmt.Println(s4.Name)
}
