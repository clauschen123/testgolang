////////////////////////////////////////////
// 统计总在线人数
////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	//"io"
	"net/http"
	"os"
	//	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type zone_report_t struct {
	Zone  int `zone`
	Count int `count`
}

type online_record_t struct {
	Ver    int
	Top    int
	TimeAt string
}

var (
	cur_count   int                   = 0
	record      online_record_t       = online_record_t{1, 0, ""}
	zone_report map[int]zone_report_t = make(map[int]zone_report_t)
)

func online_stat() {

	Init()

	http.HandleFunc("/report", ReportHandler)
	http.HandleFunc("/query", QueryHandler)
	http.ListenAndServe(":8881", nil)
}

func LoadStat() {

	f, err := os.OpenFile("online_stat.db", os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("open online_stat failed!")
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	input.Scan()

	fields := strings.Fields(input.Text())
	if len(fields) == 3 {
		record.Ver, _ = strconv.Atoi(fields[0])
		record.Top, _ = strconv.Atoi(fields[1])
		record.TimeAt = fields[2]
	}

	fmt.Println("", record)
}

func SaveStat() {

	f, err := os.OpenFile("online_stat.db", os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		//panic("open online_stat failed!")
		return
	}
	defer f.Close()

	input := bufio.NewWriter(f)

	if _, err := fmt.Fprintf(input, "%d %d %s", record.Ver, record.Top, record.TimeAt); err != nil {
		fmt.Println(err)
	}
	input.Flush()

	//f.WriteString("1 0 1466057111")

}

func Init() {
	LoadStat()
}

func CalcuTotalOnline() int {

	var total int = 0
	for _, v := range zone_report {
		total += v.Count
	}
	return total
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	sid, err := strconv.Atoi(query.Get("sid"))
	if err != nil {
		fmt.Println("sid fail")
		w.Write([]byte("Sid fail"))
		return
	}

	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		fmt.Println("count fail")
		w.Write([]byte("Count fail"))
		return
	}

	zone_report[sid] = zone_report_t{sid, count}
	cur_count = CalcuTotalOnline()

	if cur_count > record.Top {
		record.Top = cur_count
		record.TimeAt = time.Now().Format("2006-01-02_15:04:05")

		SaveStat()
	}

	w.Write([]byte("SUCC"))
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(strconv.Itoa(cur_count)))
	w.Write([]byte(" ["))
	w.Write([]byte(strconv.Itoa(record.Top)))
	w.Write([]byte(" at "))
	w.Write([]byte(record.TimeAt))
	w.Write([]byte(" ]"))
}

func toString(a interface{}) string {

	if v, p := a.(int); p {
		return strconv.Itoa(v)
	}

	if v, p := a.(float64); p {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	if v, p := a.(float32); p {
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}

	if v, p := a.(int16); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(int32); p {
		return strconv.Itoa(int(v))
	}
	return ""
}
