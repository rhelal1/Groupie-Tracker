package groupie

import (
	// "fmt"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./template/style.css")
		return

	} else if r.URL.Path == "/home.html" || r.URL.Path == "/" {
		http.ServeFile(w, r, "./template/home.html")
		return

	} else if r.URL.Path == "/aboutus.html" {
		http.ServeFile(w, r, "./template/aboutus.html")
		return

	} else if r.URL.Path == "/index.html" {
		IndexHandler(w, r)
		return

	} else if r.URL.Path == "/submit" {
		SubmitHandler(w, r)
		return

	} else if r.URL.Path == "/search" {
		SearchHandler(w, r)
		return

	} else if r.URL.Path == "/suggestions" {
		SuggHandler(w, r)
		return
	} else if r.URL.Path == "/filter" {
		filterHandler(w, r)
		return

	}  else {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, "404 page not found")
		return
	}

}

func ErrorHandler(w http.ResponseWriter, r *http.Request, s string) {
	indexTemplate, _ := template.ParseFiles("./template/ErrorPage.html")
	errorPage := ErrorPage{
		Message: s,
	}
	indexTemplate.Execute(w, errorPage)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, "400 Bad Request")
		return
	}

	indexTemplate, err := template.ParseFiles("./template/index.html")

	creationDateMin := r.FormValue("creationDateMin")
	creationDateMax := r.FormValue("creationDateMax")

	firstAlbumMin := r.FormValue("firstAlbumMin")
	firstAlbumMax := r.FormValue("firstAlbumMax")

	membersNum := r.Form["checkbox[]"]

	locationsRange := r.FormValue("locFlter")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
		return
	}

	Artist := Filters(creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax, locationsRange, membersNum)

	if len(Artist) == 0 {
		ErrorHandler(w, r, "Sorry, No Result ☹!")
		return
	}

	pageData := PageData{
		All:            Artist,
		LocFLT:         LoctionList,
		NumberOfMember: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
	}

	err = indexTemplate.Execute(w, pageData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
	}
}
