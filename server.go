package main

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

//Page - A simple structure representing web content to return
type Page struct {
	Body []byte
}

func main() {
	http.HandleFunc("/", pageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadPage(title string) (*Page, error) {
	filename := "checky/" + title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	contentType := mime.TypeByExtension(filepath.Ext(title))
	w.Header().Add("Content-type", contentType)
	w.Write(p.Body)
}
