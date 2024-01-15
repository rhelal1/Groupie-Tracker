package groupie

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, "400 Bad Request")
		return
	}

	indexTemplate, err := template.ParseFiles("./template/index.html")
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
		return
	}

	artists := Search(r.FormValue("searchValue"))

	if len(artists) == 0 {
		ErrorHandler(w, r, "Sorry, No Result ☹!")
		return
	}

	pageData := PageData{
		All: artists,
		LocFLT:         LoctionList,
		NumberOfMember: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
	}
	
	err = indexTemplate.Execute(w, pageData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, "500 Internal Server Error")
	}
}

func SuggHandler(w http.ResponseWriter, r *http.Request) {
	suggestions(r.URL.Query().Get("query"))

	// Convert the suggestions slice to JSON
	suggestionsJSON, err := json.Marshal(sugges)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(suggestionsJSON)
}
