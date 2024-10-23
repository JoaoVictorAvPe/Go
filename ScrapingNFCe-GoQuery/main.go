package main

import (
	"app/utils"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	namesSlice, statusSlice, valuesSlice := utils.GetInformationsFromSefaz()

	sefazSlice := utils.MakeSefazSlice(namesSlice, statusSlice, valuesSlice)
	sefazSlice = utils.SanitazeSefazSlice(sefazSlice)


	// for _, sefaz := range sefazSlice {
	// 	fmt.Println(sefaz.Name)
	// 	fmt.Println(sefaz.Status)
	// 	fmt.Println(sefaz.Media)
	// 	fmt.Println("=====================")
	// }

	jsonData, err := json.Marshal(sefazSlice)
	if err != nil {
		log.Fatal("Failed to convert to JSON")
	}

	fmt.Println(string(jsonData))

}
