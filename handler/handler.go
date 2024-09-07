package API

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "API/struct"
)

func handle(url string, data []api.Data) []api.Data {
	all, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return data
	}

	defer all.Body.Close()

	if err := json.NewDecoder(all.Body).Decode(&data); err != nil {
		fmt.Println("error11", err)
		return data
	}
	return data
}

func ahandel(url1 string, data any, id string, w http.ResponseWriter, r *http.Request) {
	respons, err := http.Get(url1)
	if err != nil {
		fmt.Println("error", err)
	}
	defer respons.Body.Close()

	if err := json.NewDecoder(respons.Body).Decode(&data); err != nil {
		fmt.Println("error", err)
	}
}
