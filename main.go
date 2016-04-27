package main

import (
	"Top50Anime/utils"
	"fmt"
)

func main() {
	fmt.Println(topanime.GetTopAnime(50))
}

// func getMyAnimeListHTMLIOReader() io.Reader {
// 	//get the html code from myanimelist
// 	response, err1 := http.Get("http://myanimelist.net/topanime.php?_location=mal_h_m")
// 	if nil == err1 {
// 		return response.Body
// 	}
// 	fmt.Println(err1)
// 	return nil
// }
//
// func getMALHTML() string {
// 	response, err := http.Get("http://myanimelist.net/topanime.php?_location=mal_h_m")
// 	if nil == err {
// 		htmlCode, _ := ioutil.ReadAll(response.Body)
// 		return string(htmlCode)
// 	}
// 	return ""
// }
//
// func parseTop5AnimeHTMLCode() {
// 	//var topTitles []string
// 	reader := getMyAnimeListHTMLIOReader() //the io.Reader
//
// 	tokenizer := html.NewTokenizer(reader) //used to read the html code token by tokenizer
// 	for {
// 		curToken := tokenizer.Next() //advance to next token
//
// 		//if reader is done return
// 		if html.ErrorToken == curToken {
// 			return
// 		}
//
// 		//if text is Top Anime, titles could be in following tokens...
// 		if "Top Anime" == string(tokenizer.Text()) {
// 			keepScanning := true
// 			for keepScanning {
// 				curToken = tokenizer.Next()
// 				//check for end of reader or need to escape this loop...
// 				if html.ErrorToken == curToken {
// 					return
// 				}
//
// 				_, hasAttr := tokenizer.TagName() //check if this token has attributes
// 				if hasAttr {
// 					hadClass := false
// 					//iterate through attributes
// 					for hasAttr {
// 						attrName, attrValue, moreAttr := tokenizer.TagAttr()
// 						hasAttr = moreAttr
// 						if "class" == string(attrName) && "rank" == string(attrValue) {
// 							curToken = tokenizer.Next()
// 							curToken = tokenizer.Next()
// 							//scan until next token is text..this should be anime titles
// 							for "Text" != curToken.String() {
// 								curToken = tokenizer.Next()
// 							}
// 							fmt.Println(string(tokenizer.Text()))
// 							hadClass = true
// 						}
// 					}
// 					if !hadClass {
// 						keepScanning = false
// 					}
// 				}
// 			}
// 		}
//
// 	}
// 	//return topTitles
// }
//
// //walks through char by char....painfully
// func parseUpToFiftyAnime(num int) string {
// 	s := "<img alt=\"Anime: "
// 	var listOfShows string
// 	//var listOfIndecies []int //for testing
// 	htmlCode := getMALHTML()
// 	curSlice := htmlCode[0:]
// 	for i := 0; i < num; i++ {
// 		//find next instance of img alt
// 		index := strings.Index(htmlCode, s)
// 		//listOfIndecies = append(listOfIndecies, index)
// 		listOfShows += subStr(htmlCode, index+17, "\""[0]) + "\n"
// 		htmlCode = string(curSlice[index+17:])
// 		curSlice = htmlCode[0:]
// 	}
// 	return listOfShows
// }
//
// func subStr(s string, startIndex int, endChar byte) string {
// 	var i int
// 	for i = startIndex; s[i] != endChar; i++ {
// 	}
// 	return string(s[startIndex:i])
// }

// func test() {
// 	s := "This is a test"
// 	i := strings.Index(s, "is a")
// 	slice := s[i:4]
// 	fmt.Println(s)
// 	fmt.Println(string(slice))
//
// }
