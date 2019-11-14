package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// const (
// 	INTEGER = "INTEGER"
// 	STRING  = "STRING"
// )

type ColumnType int
type SortType int

const (
	INTEGER ColumnType = iota
	STRING  ColumnType = iota
)
const (
	ASCENDING  SortType = iota
	DESCENDING SortType = iota
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

// GetRowToJSON This func will return the data from the desired row.
// You can also get all the headers, by sending 1 as row number
func GetRowToJSON(csvPath string, rowNumber int) [][]byte {
	lines := readCSV(csvPath)
	if rowNumber > 0 && rowNumber < len(lines) {
		byteArray := [][]byte{}
		jsonMap := map[string]string{}
		for j := 0; j < len(lines[0]); j++ {
			jsonMap[lines[0][j]] = lines[rowNumber-1][j]
		}
		json, err := json.Marshal(jsonMap)
		if err != nil {
			log.Println("Marshal error : ", err)
			return nil
		}
		byteArray = append(byteArray, json)
		return byteArray
	}
	log.Println("Row number should not be greater than total rows and less than 1")
	return nil

}

func GetColumnToJSON(csvPath string, columnNumber int) [][]byte {
	lines := readCSV(csvPath)
	totalNumOfColumns := 0
	if len(lines) > 0 {
		totalNumOfColumns = len(lines[0])
	}

	if columnNumber > 0 && columnNumber <= totalNumOfColumns {

		byteArray := [][]byte{}
		c := columnNumber - 1
		for i := 0; i < len(lines); i++ {
			if i == 0 {
				continue
			}
			jsonMap := map[string]string{}

			jsonMap[lines[0][c]] = lines[i][c]

			json, err := json.Marshal(jsonMap)
			if err != nil {
				log.Println("Marshal error : ", err)
				return nil
			}
			byteArray = append(byteArray, json)

		}
		return byteArray
	}
	log.Println("Column number should be less than total number of columns and greater than 0")
	return nil
}

func GetJSONByExcludingColumns(csvPath string, columnNumbers ...int) [][]byte {
	lines := readCSV(csvPath)
	byteArray := [][]byte{}
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			continue
		}
		jsonMap := map[string]string{}
		for j := 0; j < len(lines[0]); j++ {
			if !contains(columnNumbers, (j + 1)) {
				jsonMap[lines[0][j]] = lines[i][j]
			}
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

func GetSortedJSON(csvPath string, sortBy string, sortType SortType, colType ColumnType) [][]byte {

	lines := readCSV(csvPath)
	byteArray := [][]byte{}
	arrOfMaps := []map[string]string{}

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			continue
		}
		jsonMap := map[string]string{}
		for j := 0; j < len(lines[0]); j++ {
			jsonMap[lines[0][j]] = lines[i][j]
		}
		arrOfMaps = append(arrOfMaps, jsonMap)

	}

	sortColumns(arrOfMaps, sortBy, sortType, colType)

	for _, val := range arrOfMaps {
		json, err := json.Marshal(val)
		if err != nil {
			log.Println("Marshal error : ", err)
			return nil
		}
		byteArray = append(byteArray, json)
	}

	return byteArray

}

func sortColumns(arrOfMaps []map[string]string, sortBy string, sortType SortType, colType ColumnType) {

	switch colType {
	case INTEGER:
		sort.Slice(arrOfMaps[:], func(i, j int) bool {
			ival, _ := strconv.Atoi(arrOfMaps[i][sortBy])
			jval, _ := strconv.Atoi(arrOfMaps[j][sortBy])
			if sortType == ASCENDING {
				return ival < jval

			} else {
				return ival > jval

			}
		})
	case STRING:
		sort.Slice(arrOfMaps[:], func(i, j int) bool {
			ival := arrOfMaps[i][sortBy]
			jval := arrOfMaps[j][sortBy]
			if sortType == ASCENDING {
				return ival < jval

			} else {
				return ival > jval

			}
		})
	}
}

func contains(arr []int, num int) bool {
	for _, index := range arr {
		if index == num {
			return true
		}
	}
	return false
}
