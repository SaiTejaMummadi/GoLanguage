package main

import (
	"net/http"
	"io"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this a Doggy doggy doggy with a dab of ranch")


}
func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name is Sai")
}
func index(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"this is the index page")
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
