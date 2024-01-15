package groupie

import (
	"slices"
	"strings"
)

func Search(s string) []Artist {
	TheArtists := artists
	var artists []Artist
	searchWord, cond := RetutrnCond(s)
	Inlocation := AppendLocation(TheLocations.Index, searchWord)

	for index, artist := range TheArtists {
		if (cond == "Date" || cond == "") &&  Check(TheDates.Index2[index].TheData, searchWord) {
			artists = append(artists, artist)

		} else if (cond == "Member" || cond == "") && Check(artist.Members, searchWord){
			artists = append(artists, artist)

		} else if (cond == "Creation Date" || cond == "") && AppendCreationDate(artist.CreationDate, searchWord){
			artists = append(artists, artist)

		} else if (cond == "Location" || cond == "") && slices.Contains(Inlocation, index) {
			artists = append(artists, artist)

		} else if (cond == "band" || cond == "") && strings.Contains(strings.ToLower(artist.Name), searchWord) {
			artists = append(artists, artist)

		} else if (cond == "First Album" || cond == "") && strings.Contains(strings.ToLower(artist.FirstAlbum), searchWord) {
			artists = append(artists, artist)
		}
	}

	return artists
}

func Check(Members []string, searchWord string) bool {
	for i := 0; i < len(Members); i++ {
		if strings.Contains(strings.ToLower(Members[i]), searchWord) {
			return true
		}
	}
	return false
}

func RetutrnCond(s string) (string, string) {
	cond := ""
	if strings.LastIndex(s, "-") != -1 {
		cond = s[strings.LastIndex(s, "-")+2:]
		s = s[:strings.LastIndex(s, "-")-1]
	}
	return strings.ToLower(s), cond
}
