package groupie

import (
	"slices"
	"fmt"
	"strings"
)

var sugges []string

func suggestions(word string) {
	word = strings.ToLower(word)
	sugges = nil
	for x, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), word) {
			sugges = append(sugges, artist.Name+" - band")
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), word) {
			sugges = append(sugges, artist.FirstAlbum+" - First Album")
		}
		AppendCreationDate(artist.CreationDate, word)
		AppendTheMembers(artist.Members, word)
		AppendTheDates(TheDates.Index2[x].TheData, word)
		AppendLocation(TheLocations.Index, word)
	}
}

func AppendTheDates(Dates []string, word string) {
	for i := 0; i < len(Dates); i++ {
		if strings.Index(strings.ToLower(Dates[i]), word) != -1 {
			sugges = append(sugges, Dates[i]+" - Date")
		}
	}
}

func AppendTheMembers(Members []string, word string) {
	for j := 0; j < len(Members); j++ {
		if strings.Contains(strings.ToLower(Members[j]), word) && !slices.Contains(sugges, Members[j]+" - Member") {
			sugges = append(sugges, Members[j]+" - Member")
		}
	}
}

func AppendCreationDate(CreationDate interface{}, word string) bool {
	j := fmt.Sprintf("%v", CreationDate) // Print the interface data
	if j == word {
		sugges = append(sugges, j+" - Creation Date")
		return true
	}
	return false
}

func AppendLocation(Locations []Index, word string) []int {
	var index []int
	for i := 0; i < len(Locations); i++ {
		flag2 := false
		for j := 0; j < len(Locations[i].TheData); j++ {
			if strings.Contains(strings.ToLower(Locations[i].TheData[j]), word) {
				if !slices.Contains(sugges, Locations[i].TheData[j]+" - Location") {
					sugges = append(sugges, Locations[i].TheData[j]+" - Location")
				}
				if !flag2 {
					index = append(index, i)
					flag2 = true
				}
			}
		}
	}
	return index
}

