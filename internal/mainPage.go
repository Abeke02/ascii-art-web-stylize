package internal

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Message string
}

func main_page(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if url != "/" {
		ErrorHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	method := r.Method
	if method != "GET" {
		ErrorHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	temp, err1 := template.ParseFiles("./ui/templates/index.html")
	if err1 != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}
