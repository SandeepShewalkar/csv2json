package main

import (
	"CSV-TO-JSON/csv2json"
	"fmt"
)

func main() {

	jsonData := csv2json.ConverToJSONArray("/home/Records.csv")

	for _, val := range jsonData {
		fmt.Println(string(val))
	}
}
