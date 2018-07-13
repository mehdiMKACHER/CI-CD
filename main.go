package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	Title string
	Body  []byte
}

func view(w http.ResponseWriter, r *http.Request) {
	pv := &page{Title: "Demo Go@K8s v8.0 para Youtube", Body: []byte("Portal Web Golang@K8s v8.0, if you can see this yell hooray! Learn K8s")}
	t, _ := template.ParseFiles("index.html")
	fmt.Println(pv.Title)
	fmt.Println(string(pv.Body))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, pv)

}

func imagenes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, r.URL.Path[1:])
	fmt.Println(r.UserAgent())
	fmt.Println("Sirviendo Contenido")
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", view)
	http.HandleFunc("/img/", imagenes)
	http.ListenAndServe(":8188", nil)
}
