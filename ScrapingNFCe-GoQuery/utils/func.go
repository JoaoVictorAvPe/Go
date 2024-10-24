package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type Sefaz struct {
	Name   string	`json:"name"`
	Status string	`json:"status"`
	Media  string	`json:"media"`
}

func GetInformationsFromSefaz() ([]string, []string, []string) {
	document := getHTMLDocument("http://www.nfce.se.gov.br/portal/ConStatusAuto?Origem=1")

	body := document.Find("body")
	if body.Length() == 0 {
		log.Println("Error, body length is zero")
		panic("Error, body length is zero")
	}

	var namesSlice, statusSlice, valuesSlice []string

	body.Contents().Each(func(i int, element *goquery.Selection) {
		nodeName := goquery.NodeName(element)

		switch nodeName {
		case "a":
			namesSlice = append(namesSlice, element.Text())
		case "img":
			alt, exists := element.Attr("alt")
			if exists {
				statusSlice = append(statusSlice, alt)
			} else {
				statusSlice = append(statusSlice, "none")
			}
		case "#text":
			valuesSlice = append(valuesSlice, element.Text())
		default:
			errorMessage := "A new tag has been identified in the document"
			log.Println(errorMessage)
			panic(errorMessage)
		}
	})

	return namesSlice, statusSlice, valuesSlice
}

func getHTMLDocument(url string) *goquery.Document {
	response, err := http.Get(url)
	if err != nil {
		errorMessage := "Faild to get NFCe web page: " + err.Error()
		log.Println(errorMessage)
		panic(errorMessage)
	}
	if response.StatusCode != 200 {
		errorMessage := fmt.Sprintf("HTTP Error %d: %s", response.StatusCode, response.Status)
		log.Println(errorMessage)
		panic(errorMessage)
	}
	defer response.Body.Close()

	utf8Reader, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		errorMessage := "Faild to convert web page to UTF-8: " + err.Error()
		log.Println(errorMessage)
		panic(errorMessage)
	}

	document, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return document
}

func MakeSefazSlice(names, status, values []string) []Sefaz {
	var SizeOfTheBiggestSlice int = len(biggestSlice(names, status, values))
	var sefazSlice []Sefaz

	for i := range SizeOfTheBiggestSlice {
		var sefaz Sefaz
		if i < len(names) {
			sefaz.Name = names[i]
		}
		if i < len(status) {
			sefaz.Status = status[i]
		}
		if i < len(values) {
			sefaz.Media = values[i]
		}

		sefazSlice = append(sefazSlice, sefaz)
	}

	return sefazSlice
}

func biggestSlice(slice1, slice2, slice3 []string) []string {
	s1 := len(slice1)
	s2 := len(slice2)
	s3 := len(slice3)

	var biggestSlice []string
	biggetSliceSize := len(biggestSlice)

	if s1 >= biggetSliceSize {
		biggestSlice = slice1
	}
	if s2 >= biggetSliceSize {
		biggestSlice = slice2
	}
	if s3 >= biggetSliceSize {
		biggestSlice = slice3
	}

	return biggestSlice
}

func SanitazeSefazSlice(sefazSlice []Sefaz) []Sefaz {
	var sanitizedSlice []Sefaz

	for i, sefaz := range sefazSlice {
		sliceSplited := strings.Split(sefaz.Media, "ms")
		sefaz.Media = sliceSplited[2] + "ms"

		if sefaz.Name == "" {
			valuesAnteriorSplited := strings.Split(sefazSlice[i-1].Media, "ms")
			sefaz.Name = valuesAnteriorSplited[len(valuesAnteriorSplited)-1]
		}

		sanitizedSlice = append(sanitizedSlice, sefaz)
	}

	return sanitizedSlice
}
