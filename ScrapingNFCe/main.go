package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("http://www.nfce.se.gov.br/portal/painelMonitor.jsp")
	if err != nil {
		panic("Faild to get NFCe web page")
	}
	if response.StatusCode != 200 {
		panic(fmt.Sprintf("HTTP Error %d: %s", response.StatusCode, response.Status))
	}
	defer response.Body.Close()


	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	htmlTable := document.Find("#tabDados")
	tableBody := htmlTable.Find("tbody")
	fmt.Println(tableBody.Text())

}