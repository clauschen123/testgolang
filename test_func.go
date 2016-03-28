package main

import "fmt"

//multi param
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func divide2(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func sum(aregs ...int) int {
	s := 0
	for _, number := range aregs {
		s += number
	}
	return s
}

//function type
type MyFuncType func(int) bool

func IsBigThan5(n int) bool {
	return n > 5
}

func Display(arr []int, f MyFuncType) {
	for _, v := range arr {
		if f(v) {
			fmt.Println(v)
		}
	}
}

//exception
func Test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	divide(5, 0)               //程序出错，中断执行
	fmt.Println("end of test") //该语句不会被执行
}

func test_func() {
	Test()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Display(arr, IsBigThan5)

	total := sum(1, 2, 3, 4)
	fmt.Println(total)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //定义一个切片
	/*将切片传 sum 时，要用...展开否则将做为一个参数处理
	  等价于 sum(1,2,3,4,5,6,7,8,9)
	*/
	total = sum(slice...)
	fmt.Println(total)
}
