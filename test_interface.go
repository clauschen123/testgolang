package main

import "fmt"

// sample 1
type IDemo interface {
	GetX() int
}
type Demo_t struct {
	x int
}

func (d Demo_t) GetX() int {
	return d.x
}

func get_data() *Demo_t {
	return &Demo_t{x: 100}
}

//指针也是Interface
func GetData() IDemo {
	return &Demo_t{x: 100}
}
func GetData2() IDemo {
	return Demo_t{x: 100}
}

type Stringer interface {
	String() string
}

// sample 2
//var void interface{} // Value provided by caller.
type void interface{}
type INT int

//func (p *INT) String() string {
//	return string("INT")
//}

func GetString(void interface{}) string {
	switch str := void.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}

	return string("haha")
}

func test_interface() {
	//-------------------------------------------------
	var n INT = 10
	var p *INT = &n

	var v void = n
	fmt.Println(v.(INT)) //assigned type

	v = p
	t, ok := v.(*INT)
	if ok {
		fmt.Printf("T value is: %q\n", t)
	} else {
		fmt.Printf("value is not a T type \n")
	}
	//	if _, err := fmt.Println(v.(INT)); err != nil {
	//	}

	//var str string = "sfdf"
	//fmt.Println(GetString(str))
	//fmt.Println(GetString(n))
	//fmt.Println(GetString(str))

	//-------------------------------------------------
	var demo IDemo = GetData()
	switch t := demo.(type) { // must be an interface type
	default:
		fmt.Printf("%T\n", t)
	}

	fmt.Printf("%T\n", demo)
}
