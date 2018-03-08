package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"
	//	_ "github.com/ziutek/mymysql/godrv"
)

type Userinfo struct {
	Uid        int `PK` //如果表的主键不是 id，那么需要加上pk 注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

func test_beedb() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.1.107:3306)/test?charset=utf8")
	//why the sample is:
	//db, err := sql.Open("mymysql", "test/xiemengjun/123456")

	if err != nil {
		panic(err)
	}
	beedb.OnDebug = true
	orm := beedb.New(db)

	add := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2012-12-02"
	_, err = orm.SetTable("userinfo").Insert(add)

	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	fmt.Println(saveone.Uid)
	err = orm.Save(&saveone)
	if err != nil {
		panic(err)
	}
	fmt.Println(saveone.Uid)
}
