package main
import(
	"fmt"
	"github.com/astaxie/goredis"
)

func test_redis(){
	var client goredis.Client
	client.Addr = "192.168.1.109:6379"
    client.Db = 1
	
	//字符串操作
	client.Set("a", []byte("clausgo"))
	val, _ := client.Get("a")
	fmt.Println(string(val))
	//client.Del("a")
	
	//list 操作
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals,_ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i,":",string(v))
	}
	//client.Del("l")
}