package API

import (
	"fmt"
	"html/template"
	"net/http"

	api "API/struct"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		PageError(w, r, http.StatusNotFound, "page not fond")
		return
	}
	url := ("https://groupietrackers.herokuapp.com/api/artists")

	var slice []api.Data

	data := handle(url, slice)

	tmp, err := template.ParseFiles("html/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmp.Execute(w, data)
}
