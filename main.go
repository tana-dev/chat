package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	ttt := t.templ.Execute(w, nil)
	fmt.Println(ttt)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "main.html"})
	http.Handle("/room", r)
	go r.run()
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte(`
	//		<html>
	//			<head>
	//				<title>tanaka</title>
	//			</head>
	//			<body>
	//			</body>
	//		</html>
	//		`))
	//	})

	//	http.Handle("/", &templateHandler{filename: "main.html"})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
