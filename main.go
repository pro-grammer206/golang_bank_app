package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var err error

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./assets"))))
	tpl, err = tpl.ParseGlob("./assets/pages/*.gohtml")
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/authlogin", loginHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "the form values are %v", r.Form)

}
