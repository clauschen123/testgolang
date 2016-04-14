package main

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

//test_json3: json2map
//把请求包定义成一个结构体
type Requestbody struct {
	req string
}

//以指针的方式传入，但在使用时却可以不用关心,result 是函数内的临时变量，作为返回值可以直接返回调用层
func (r *Requestbody) Json2map() (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func test_json3() {
	//json转map
	var r Requestbody
	r.req = `{"name": "xym","sex": "male"}`
	if req2map, err := r.Json2map(); err == nil {
		fmt.Println(req2map["name"])
		fmt.Println(req2map)
	} else {
		fmt.Println(err)
	}
}

//test_json2: map2json
type ChannelOperator struct {
	Oid      string
	Aid      string
	Name     string
	Isonline bool
	//	Msgchan  chan string `json:"-"`
}

type ChannelOperators struct {
	Op []ChannelOperator
}

func test_json2() {
	fmt.Println("Hello, map2json")
	mm := make(map[string]*ChannelOperator)
	mm["123"] = &ChannelOperator{Oid: "12312"}
	mm["5678"] = &ChannelOperator{Oid: "12312"}

	b, _ := json.Marshal(mm)
	fmt.Println(string(b))

	c := ChannelOperators{}
	c.Op = append(c.Op, ChannelOperator{Oid: "12312"})

	x, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))
}

//成员变量必须大写，否则UnMarshal后没写进去
//改成tag名怎么不行
type ZoneInfo struct {
	Zone  string `zone`
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
