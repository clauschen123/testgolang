package main

import "fmt"

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

func test_base() {
	fmt.Printf("0x%x\n", KB)
	fmt.Printf("0x%x\n", MB)
	fmt.Printf("0x%x\n", GB)
	fmt.Printf("0x%x\n", TB)
}
