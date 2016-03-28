package main

import (
	"fmt"
	"reflect"
)

func test_tag() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	//取 tag 数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	/*可以你 JSON 一样，取 TAG 里的数据，注意，设置时，两个之间无逗
	  号,键名无引号*/
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}
}

//call fucn
func test_call() {
	s := Student{Name: "abc", Age: 19}
	rt := reflect.TypeOf(&s)
	rv := reflect.ValueOf(&s)
	fmt.Println("typeof 调用函数")
	rtm, ok := rt.MethodByName("PrintName")
	if ok {
		var parm []reflect.Value
		//函数默认第一个参数是结构体本身即*Student
		parm = append(parm, rv)
		rtm.Func.Call(parm)
	}
	//valueof 调用函数
	fmt.Println("valueof 调用函数")
	rvm := rv.MethodByName("GetAge")
	//用 valueof 调用函数时不需要把 Struct 本身做为参数传递过去
	ret := rvm.Call(nil)
	//显示返回值
	fmt.Println("返回值:", ret)
	ShowSlice(ret)
}

func ShowSlice(s []reflect.Value) {
	if s != nil && len(s) > 0 {
		for _, v := range s {
			fmt.Println(v.Interface())
		}
	}
}
func test_reflect() {

	test_tag()
	return

	s := Student{Name: "abc", Age: 19}
	rt := reflect.TypeOf(s)
	newrt := reflect.TypeOf(rt)

	//判断是否指针类型，如果是，取指针所指向的元素的类型
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	//输出类型所在的包的路径
	fmt.Println(rt.PkgPath())
	fmt.Println(newrt.PkgPath())

	//反射取所有字段
	fmt.Println(rt.Name(), "共有", rt.NumField(), "个字段")
	for i, j := 0, rt.NumField(); i < j; i++ {
		rtField := rt.Field(i)
		fmt.Println(rtField.Name)
	}

	/*因为我们的函数定义是在*Student 类型上的，所以这里转换为指针类型， 否
	  则反射会取不到函数*/
	rt = reflect.PtrTo(rt)
	//反射取所有函数
	fmt.Println(rt.Name(), "共有", rt.NumMethod(), "个函数")
	for i, j := 0, rt.NumMethod(); i < j; i++ {
		mt := rt.Method(i)
		fmt.Println(mt.Name)
		//输入参数的数量
		numIn := mt.Type.NumIn()
		//输出参数的数量
		numOut := mt.Type.NumOut()
		//输出输入参数
		if numIn > 0 {
			fmt.Println("\t 共", numIn, "个输入参数")
			for k := 0; k < numIn; k++ {
				in := mt.Type.In(k)
				fmt.Println("\t", in.Name(), "\t", in.Kind())
			}
		}
		//输出输出参数
		if numOut > 0 {
			fmt.Println("\t 共", numOut, "个输出参数")
			for k := 0; k < numOut; k++ {
				out := mt.Type.Out(k)
				fmt.Println("\t", out.Name(), "\t", out.Kind())
			}
		}
	}

	//TypeOf 只能取到字段名，字段类型，取不到字段值；要取字段值，需要用 ValueOf。
	rv := reflect.ValueOf(s)
	//判断是否指针类型，如果是，取指针所指向的元素的类型
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rvField := rv.FieldByName("Name") //取 Name 字段的值
	fmt.Println(rvField.String())

	rvField = rv.FieldByName("Age") //取 Age 字段的值
	fmt.Println(rvField.Interface())
	fmt.Println(rvField.Int())

	//SetString(),SetBool(),SetInt()等用来设置反射对像的值。反射时必须是
	//对指针进行反射，因为值类型的参数，在函数内被改变时不会改外函数外的值
	rv = reflect.ValueOf(&s)
	//判断是否指针类型，如果是，取指针所指向的元素的类型
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rvField = rv.FieldByName("Name") //取 Name 字段的值
	fmt.Println(rvField.String())
	rvField.SetString("已改名")
	fmt.Println(s.Name) //输出已改名
	fmt.Print(s)

}
