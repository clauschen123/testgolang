package main

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

//成员变量必须大写，否则UnMarshal后没写进去
//改成tag名怎么不行
type ZoneInfo struct {
	Zone  uint16 `zone`
	Total string `total`
}

func test_json() {
	s := &Student{Name: "张三", Age: 19}
	//将 s 编码为 json
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buf))
	//将 json 字符串转换成 Student 对像
	var s1 Student
	json.Unmarshal(buf, &s1)
	fmt.Println(s1)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	b = []byte(`{
			"zone" : 37,
			"uid" : "166633186212710420",
			"add" : true,
			"total" : "1"
			}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	fmt.Println(f)

	fmt.Println("....................")
	b1 := []byte(`{
			"zone" : 37,
			"uid" : "166633186212710420",
			"add" : true,
			"total" : "1"
			}`)
	var zi ZoneInfo
	err = json.Unmarshal(b1, &zi)
	fmt.Println(zi)

}

func test_xml() {
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
