package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := template.ParseFiles("./template/index.html")

	if !flag || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
		return
	}

	pageData := PageData{
		All:            artists,
		LocFLT:         LoctionList,
		NumberOfMember: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
	}

	err = indexTemplate.Execute(w, pageData)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
	}
}
