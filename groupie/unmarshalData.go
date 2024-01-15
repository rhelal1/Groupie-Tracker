package groupie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

func readThefile(url string) []byte {
	var body []byte
	var err error
	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		flag = false
		fmt.Print(err)
	} else {
		defer response.Body.Close()

		// Read response body
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
		flag = false
		fmt.Print(err)
		}
	}

	// Print raw JSON response
	return body
	// Parse JSON
}

func UnmarshalData() {
	flag = true
	//for artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	body := readThefile(url)
	err := json.Unmarshal(body, &artists)
	if err != nil {
		flag = false
		fmt.Print(err)
	}

	//for the locations
	url = "https://groupietrackers.herokuapp.com/api/locations"
	data := readThefile(url)
	err = json.Unmarshal(data, &TheLocations)
	if err != nil {
		flag = false
		fmt.Print(err)
	}

	for z := 0; z < len(TheLocations.Index); z++ {
		for i := 0; i < len(TheLocations.Index[z].TheData); i++ {
			TheLocations.Index[z].TheData[i] = strings.ReplaceAll(TheLocations.Index[z].TheData[i], "_", " ")
		}
	}

	// for the dates
	url = "https://groupietrackers.herokuapp.com/api/dates"
	Datees := readThefile(url)
	err = json.Unmarshal(Datees, &TheDates)
	if err != nil {
		flag = false
		fmt.Print(err)
	}

	// for the relations
	url = "https://groupietrackers.herokuapp.com/api/relation"
	relations := readThefile(url)
	err = json.Unmarshal(relations, &TheRelations)
	if err != nil {
		fmt.Print(err)
	}
	if flag {
		artists[20].Image = "https://cdns-images.dzcdn.net/images/artist/94abb0f5039ec687e2f1413c96e64d68/500x500.jpg"
		LocationFilter("")
	}
}
