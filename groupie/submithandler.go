package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if !flag {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
		return
	}

	valueStr := r.URL.Query().Get("value")
	value, _ := strconv.Atoi(valueStr)
	indexTemplate, err := template.ParseFiles("./template/artist.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
		return
	}

	if valueStr == "" || value > 52 || value < 1 {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, "400 Bad Request")
		return
	}

	pageDataArtice := PageDataArtice{
		All:                    artists[value-1],
		MergeDatesAndLocations: MergeDatesAndLocations(value - 1),
	}

	err = indexTemplate.Execute(w, pageDataArtice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
	}
}
