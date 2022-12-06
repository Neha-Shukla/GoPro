package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
type Student struct {
    Name string
    College string
    RollNumber int64
}
	 student := Student{
        Name:       "GB",
        College:    "GolangBlogs",
        RollNumber: 1,
    }

	partials := []string{
		"./templates/base.layout.gohtml",
		"./templates/header.partial.gohtml",
		"./templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
