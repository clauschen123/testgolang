package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func run3() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

//execute a shell file
func run2() {
	cmd := exec.Command("/bin/sh", "./exec.sh")

	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	fmt.Println(string(bytes))
}

func run() {
	cmd := exec.Command("cmd.exe", "-c", "ping 127.0.0.1")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
}

func test_shell() {
	run3()
	return

	go run()
	time.Sleep(1e9)

	cmd := exec.Command("/bin/sh", "-c", `ps -ef | grep -v "grep" | grep "ping"`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Printf("stdout: %s", bytes)
}
