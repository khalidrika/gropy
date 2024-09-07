package API

import (
	"html/template"
	"net/http"
)

type error struct {
	code int
	msg  string
}

func PageError(w http.ResponseWriter, r *http.Request, c int, m string) {
	misagerr := &error{code: c, msg: m}
	tep, err := template.ParseFiles("html/err.html")
	if err != nil {
		http.Error(w, "enternal server error", 404)
		return
	}
	w.WriteHeader(c)
	err = tep.Execute(w, misagerr)
	if err != nil {
		http.Error(w, "internal server error", 404)
	}
}