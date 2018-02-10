package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this a Doggy doggy doggy with a dab of ranch")
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "dog.gohtml", "dog")
	if err != nil {
		log.Fatalln(err)
	}

}
func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name is Sai")
}
func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is the index page")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
