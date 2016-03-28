package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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
	http.HandleFunc("/test", HandleRequest)
	http.ListenAndServe(":8888", nil)
}
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>第一个 WEB 应用</h1>"))
	w.Write([]byte(r.URL.Path))
}

//sample3: a simple QR server
func test_QRserver() {
	var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18
	var templ = template.Must(template.New("qr").Parse(templateStr))

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
