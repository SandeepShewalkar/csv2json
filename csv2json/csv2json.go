package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func readCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil
	}
	return lines
}

// ConverToJSONArray will read the data from csv and convert it to json array
func ConverToJSONArray(csvPath string) [][]byte {
	lines := readCSV(csvPath)
	byteArray := [][]byte{}
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			continue
		}
		jsonMap := map[string]string{}
		for j := 0; j < len(lines[0]); j++ {
			jsonMap[lines[0][j]] = lines[i][j]
		}
		json, err := json.Marshal(jsonMap)
		if err != nil {
			log.Println("Marshal error : ", err)
			return nil
		}
		byteArray = append(byteArray, json)

	}
	return byteArray
}
