package main

import (
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"os"
)

func test_serilize_xml() {
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	//若class是private则这里初始化失败
	s := &Student{"张三", 19, "二班"}
	//创建 encode 对像
	encoder := xml.NewEncoder(f)
	//将 s 序列化到文件中
	encoder.Encode(s)
	//重置文件指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := xml.NewDecoder(f)

	var s1 Student
	//从文件中反序列化成对像
	decoder.Decode(&s1)
	fmt.Println(s1)
}

func test_gob() {
	s := &Student{"张三", 19, "erban"}
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	//创建 Encoder 对像
	encode := gob.NewEncoder(f)
	//将 s 序列化到 f 文件中
	encode.Encode(s)
	//重置文件指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := gob.NewDecoder(f)

	var s1 Student
	//反序列化对像
	decoder.Decode(&s1)
	fmt.Println(&s1)
}
