package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func test_sqlite() {

}

/***
通情况下我们用 Query,或 Exec 就可以了。
但对行重复性的操，比如，循环向数据库中插入 10000 条数据，这时就要用 Prepare 了，可
以提高程序的性能
*/
func test_mysql2() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.1.107:3306)/dev_ctj_gamedb?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var smt *sql.Stmt
	smt, err = db.Prepare("insert into person(myname,age,IsBoy) values(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("开始插入数据....", time.Now())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		_, err = smt.Exec(fmt.Sprintf("张%d", r.Int()), r.Intn(50),
			r.Intn(100)%2)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("数据插入完成!", time.Now())
}
func test_mysql() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.1.107:3306)/dev_ctj_gamedb?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var result sql.Result
	//向数据库中插入一条数据
	result, err = db.Exec("insert into person(myname,age,IsBoy) values(?,?,?)", "张三", 19, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	lastId, _ := result.LastInsertId()

	fmt.Println("新插入的数据 ID 为", lastId)
	var row *sql.Row
	//返回一行数据
	row = db.QueryRow("select * from person")
	var name string

	var id, age int
	var isBoy bool
	//取数据进行显示
	err = row.Scan(&id, &name, &age, &isBoy)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	//再插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy) values(?,?,?)", "王红", 18, false)
	fmt.Println("=====================")
	var rows *sql.Rows
	rows, err = db.Query("select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name string
		var id, age int
		var isBoy bool
		rows.Scan(&id, &name, &age, &isBoy)
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	}
	rows.Close()
	// 最后，清空表
	db.Exec("truncate table person")
}
