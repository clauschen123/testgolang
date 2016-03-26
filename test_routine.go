package main

import (
	"fmt"
	"runtime"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Print("Hello ")
		runtime.Gosched()
	}
}
func SayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("World!")
		runtime.Gosched()
	}
}

func test_1() {
	runtime.GOMAXPROCS(4)
	//n := runtime.GOMAXPROCS(runtime.NumCPU())

	go SayHello()
	go SayWorld()
	time.Sleep(5 * time.Second)

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

func test() {
	defer func() {
		fmt.Println("in defer!")
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 5 {
			runtime.Goexit()
		}
	}
}

func test_routine() {
	go test()
	var str string
	fmt.Scan(&str)
}
