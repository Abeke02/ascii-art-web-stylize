package internal

import (
	"html/template"
	"net/http"

	"main.go/ascii-art"
)

func asciiArt_page(w http.ResponseWriter, r *http.Request) {
	temp, err1 := template.ParseFiles("./ui/templates/index.html")
	if err1 != nil {
		ErrorHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	url := r.URL.String()
	if url != "/asciiArt" {
		ErrorHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	t := r.FormValue("text")
	s := r.FormValue("style")

	method := r.Method
	if method != "POST" {
		ErrorHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := errorCheck(w, t, s)
	if err != nil {
		return
	}

	finalArt, err2 := ascii.Art(t, s)
	if err2 != nil {
		ErrorHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := ViewData{
		Message: finalArt,
	}

	temp.Execute(w, data)
}
