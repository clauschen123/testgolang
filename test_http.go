package main

import (
	//	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

//---------------------------------------
//sample3: file upload
func test_upload() {
	http.HandleFunc("/", FileUpload)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(handler.Header)
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename,
			os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		size, err := io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "上传文件的大小为: %d", size)
		return
	}
	// 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	html := `
		<form enctype="multipart/form-data" action="/" method="POST">
		请选择要上传的文件: <input name="file" type="file" /><br/>
		<input type="submit" value="Upload File" />
		</form>
		`
	io.WriteString(w, html)
}

//sample2
func test_http2() {
	hadler := &HttpHandler{}
	http.ListenAndServe(":8888", hadler)
}

type HttpHandler struct {
}

func (this *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>在 ServeHTTP 里</h1>"))
	w.Write([]byte(r.URL.Path))
}

//sample1
func test_http1() {
	http.HandleFunc("/test1", HandleRequest1)
	http.HandleFunc("/test2", HandleRequest2)
	http.HandleFunc("/test3", HandleRequest3)
	http.HandleFunc("/test4", HandleRequest4)
	http.HandleFunc("/test5", HandleRequest5)
	http.HandleFunc("/test6", HandleRequest6)
	http.HandleFunc("/test100", HandleRequest100)
	http.HandleFunc("/test101", HandleRequest101)
	http.ListenAndServe(":8881", nil)
}

func HandleRequest101(w http.ResponseWriter, r *http.Request) {
	//	g_ZI[1] = ZoneInfo{"1", "zone1"}
	//	g_ZI[2] = ZoneInfo{"2", "zone2"}

	//	buf, err := json.Marshal(g_ZI)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	fmt.Println(string(buf))
	//	w.Write(buf)
}

func HandleRequest100(w http.ResponseWriter, r *http.Request) {
	//	data := r.URL.Query()

	//	jsonStr := data.Get("json")
	//	var zi ZoneInfo
	//	if len(jsonStr) > 0 {
	//		err := json.Unmarshal([]byte(jsonStr), &zi)
	//		if err != nil {
	//			fmt.Println("Unmarshal result:", err)
	//		}
	//	}
	//g_ZI[zi.Zone] = zi
	//	g_ZI[1] = zi

	w.Write([]byte("SUCC"))
}

func HandleRequest1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>第一个 WEB 应用</h1>"))
	w.Write([]byte(r.URL.Path))
}
func HandleRequest2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>第一个 WEB 应用</h1>"))
	w.Write([]byte(r.URL.Path))
}
func HandleRequest3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"name":"claus","age": 26}`))
	w.Write([]byte(r.URL.Path))
}

func HandleRequest4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>URL 参数</h1>"))
	w.Write([]byte(fmt.Sprintf("%v", r.URL.Query())))

	fmt.Println(r.URL.Query())
}

func HandleRequest5(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", " text/html;charset=utf-8")
	if r.Method == "POST" {
		r.ParseForm()
		/*username 有两个值，默认取的是第一个的*/
		w.Write([]byte("用户名：" + r.FormValue("username") + "<br/>"))
		w.Write([]byte("<hr/>"))
		names := r.Form["username"]
		w.Write([]byte("username 有两个：" + fmt.Sprintf("%v", names)))
		w.Write([]byte("<hr/>r.Form 的 内 容 ： " + fmt.Sprintf("%v", r.Form)))
		w.Write([]byte("<hr/>r.PostForm 的内容：" + fmt.Sprintf("%v", r.PostForm)))
	} else {
		strBody := `<form action="` + r.URL.RequestURI() +
			`"method="post">
			用户名：<input name="username" type="text" /><br />
			用户名：<input name="username" type="text" /><br />
			<input id="Submit1" type="submit" value="submit" />
			</form>`
		w.Write([]byte(strBody))
		r.ParseForm()
	}
}
func HandleRequest6(w http.ResponseWriter, r *http.Request) {
	strTemplate := "<div><h3>欢迎光临！</h3></div>"
	t := template.New("test")
	//解析模板
	t, err := t.Parse(strTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

//sample3: a simple QR server
var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18
var templ = template.Must(template.New("qr").Parse(templateStr))

func test_QRserver() {

	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	fmt.Println("0-30249")
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
