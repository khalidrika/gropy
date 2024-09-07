package API

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	api "API/struct"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	url1 := ("https://groupietrackers.herokuapp.com/api/")
	var data api.Artist

	id := strings.Trim(r.URL.Path, "/artist/")
	if len(id) > 100 {
		PageError(w, r, 404, "page not fond")
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil || num < 0 || num > 52 {
		PageError(w, r, 404, "page not fond")
		return
	}

	ahandel(url1+"artists/", &data.Information, strconv.Itoa(num), w, r)
	ahandel(url1+"locations/", &data.Location, strconv.Itoa(num), w, r)
	ahandel(url1+"dates/", &data.Dates, strconv.Itoa(num), w, r)
	ahandel(url1+"relation/", &data.Rolation, strconv.Itoa(num), w, r)

	tmp, err := template.ParseFiles("html/artist.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmp.Execute(w, data)
	// fmt.Println(r.FormValue("id"))
	// fmt.Println(r.URL.Query().Get("id"))
	// id := r.URL.Path[len("/artist/"):]
	// fmt.Println(id)
}
