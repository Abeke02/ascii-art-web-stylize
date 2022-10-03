package internal

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type ErrorData struct {
	ErrStatusInt    int
	ErrStatusString string
}

func errorCheck(w http.ResponseWriter, t, s string) error {
	arr := strings.Fields(t)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > 127 || arr[i][j] < 32 {
				ErrorHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return errors.New("error")
			}
		}
	}

	if t == "" {
		temp, err1 := template.ParseFiles("./ui/templates/index.html")
		if err1 != nil {
			ErrorHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return errors.New("error")
		}
		data := ViewData{
			Message: "Enter the text",
		}
		temp.Execute(w, data)
		return errors.New("error")
	}

	return nil
}

func ErrorHandle(w http.ResponseWriter, Errortxt string, Errorstatus int) {
	t, _ := template.ParseFiles("./ui/templates/error.html")
	data := ErrorData{
		ErrStatusInt:    Errorstatus,
		ErrStatusString: Errortxt,
	}
	w.WriteHeader(Errorstatus)
	t.Execute(w, data)
}
