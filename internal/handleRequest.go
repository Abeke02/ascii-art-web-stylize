package internal

import (
	"fmt"
	"net/http"
)

func HandleRequest() {
	fmt.Println("click on the link", "'http://localhost:8070/'")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	http.Handle("/ui/image/", http.StripPrefix("/ui/image/", http.FileServer(http.Dir("./ui/image"))))
	http.HandleFunc("/", main_page)
	http.HandleFunc("/asciiArt", asciiArt_page)
	http.ListenAndServe(":8070", nil)
}
