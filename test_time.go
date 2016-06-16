package main

import (
	"fmt"
	"time"
)

func test_time() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02_15:04:05"))
	fmt.Println(now.Format("2006-01-01 15:04:05"))
	fmt.Println(now.UnixNano())
	fmt.Println(now)
}
