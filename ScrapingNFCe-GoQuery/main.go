package main

import (
	"app/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("/script/status_NFCe/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open the log file: ", err)
	}
	defer file.Close()
	log.SetOutput(file)

	namesSlice, statusSlice, valuesSlice := utils.GetInformationsFromSefaz()

	sefazSlice := utils.MakeSefazSlice(namesSlice, statusSlice, valuesSlice)
	sefazSlice = utils.SanitazeSefazSlice(sefazSlice)

	jsonData, err := json.Marshal(sefazSlice)
	if err != nil {
		log.Fatal("Failed to convert to JSON")
	}

	fmt.Println(string(jsonData))

}
