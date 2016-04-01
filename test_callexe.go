package main

import (
	"fmt"
	"os/exec"
)

func test_callexe() {
	callEXE2()
	var str string
	fmt.Scan(&str)
}

func callEXE1() {
	arg := []string{}
	cmd := exec.Command("test", arg...)
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
}

func callEXE2() {
	arg := []string{}
	cmd := exec.Command("test", arg...)

	//会向 cmd.Stdout写入信息
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v", string(output) /*, cmd.Stdout, cmd.Stderr*/)
}
