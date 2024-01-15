package groupie

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Filters(creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax, locationsRange string, membersNum []string) []Artist {
	var Artists []Artist
	for index, artist := range artists {
		find := true
		if !CheckDate(fmt.Sprintf("%v", artist.CreationDate), creationDateMin, creationDateMax) || !CheckDate(artist.FirstAlbum[6:], firstAlbumMin, firstAlbumMax) {
			continue
		}

		if len(membersNum) != 0 {
			find = false
			for _, number := range membersNum {
				if number == strconv.Itoa(len(artist.Members)) {
					find = true
					break
				}
			}
		}

		if !find || (locationsRange != "" && !slices.Contains(AppendLocation(TheLocations.Index, strings.ToLower(strings.ReplaceAll(locationsRange, ", ", "-"))), index)) {
			continue
		}

		Artists = append(Artists, artist)
	}

	return Artists
}

func CheckDate(filterDate, minDate, maxDate string) bool {
	filter, _ := strconv.Atoi(filterDate)
	min, err1 := strconv.Atoi(minDate)
	max, err2 := strconv.Atoi(maxDate)

	if (err1 == nil && min > filter) || (err2 == nil && filter > max) {
		return false
	}

	return true
}

func LocationFilter(s string) {
	LoctionList = append(LoctionList, "")

	for range artists {
		Locations := TheLocations.Index
		for i := 0; i < len(Locations); i++ {
			for j := 0; j < len(Locations[i].TheData); j++ {
				if !slices.Contains(LoctionList, strings.ReplaceAll(strings.Title(Locations[i].TheData[j]), "-", ", ")) {
					LoctionList = append(LoctionList, strings.ReplaceAll(strings.Title(Locations[i].TheData[j]), "-", ", "))
				}
				index := strings.LastIndex(Locations[i].TheData[j], "-")

				if !slices.Contains(LoctionList, strings.Title(Locations[i].TheData[j][index+1:])) {
					LoctionList = append(LoctionList, strings.Title(Locations[i].TheData[j][index+1:]))
				}
			}
		}
	}

	sort.SliceStable(LoctionList, func(i, j int) bool {

		country1, city1 := extractCountryAndCity(LoctionList[i])
		country2, city2 := extractCountryAndCity(LoctionList[j])

		// Compare countries first
		if country1 != country2 {
			return country1 < country2
		}

		// If countries are the same, compare cities
		if city1 != city2 {
			return city1 < city2
		}

		// If both country and city are the same, maintain the original order
		return i < j
	})

}

func extractCountryAndCity(s string) (string, string) {
	parts := strings.SplitN(s, ", ", 2)

	if len(parts) == 2 {
		return parts[1], parts[0]
	}

	return s, ""
}
