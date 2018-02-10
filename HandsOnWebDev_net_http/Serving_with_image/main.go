package main

import (
	"net/http"

	"html/template"
	"io"
	"log"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FOO Rannnnn!!!")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", bar)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
func bar(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "maxresdefault.jpg")
}
