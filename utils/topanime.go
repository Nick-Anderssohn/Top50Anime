package topanime

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//numAnime must be less than 50
func GetTopAnime(numAnime int) string {
	if numAnime < 51 && numAnime > 0 {
		return parseUpToFiftyAnime(numAnime)
	}
	return "The number of anime requested must be greater than 0 and less than 51"
}

func getMALHTML() string {
	response, err := http.Get("http://myanimelist.net/topanime.php?_location=mal_h_m")
	if nil == err {
		htmlCode, _ := ioutil.ReadAll(response.Body)
		return string(htmlCode)
	}
	return ""
}

//walks through char by char....painfully
func parseUpToFiftyAnime(num int) string {
	s := "<img alt=\"Anime: "
	var listOfShows string
	//var listOfIndecies []int //for testing
	htmlCode := getMALHTML()
	curSlice := htmlCode[0:]
	for i := 0; i < num; i++ {
		//find next instance of img alt
		index := strings.Index(htmlCode, s)
		//listOfIndecies = append(listOfIndecies, index)
		listOfShows += subStr(htmlCode, index+17, "\""[0]) + "\n"
		htmlCode = string(curSlice[index+17:])
		curSlice = htmlCode[0:]
	}
	return listOfShows
}

func subStr(s string, startIndex int, endChar byte) string {
	var i int
	for i = startIndex; s[i] != endChar; i++ {
	}
	return string(s[startIndex:i])
}
