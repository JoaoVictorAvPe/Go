package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://globo.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h1").Each(func(i int, s *goquery.Selection){
		title := s.Text()
		fmt.Println("Titulo encontrado:", title)
	})
}