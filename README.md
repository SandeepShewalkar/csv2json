# csv2json


Simplified way to read csv files and operate over data.
This package is helpful to :
	1. covert csv file data to JSON array.
	2. Get data from desired column
	3. Get data from desired row
	4. Get data by excluding columns
	5. Column wise data sorting


# Installation

go get github.com/SandeepShewalkar/csv2json


# Sample Usage


	ConverToJSONArray: 
	-----------------
	
	csv2json.ConverToJSONArray("/home/Records.csv")

	This function gets all data in the form of JSON array



	GetRowToJSON:
	-------------

	csv2json.GetRowToJSON("/home/Records.csv", 14)

	This function gets data from row number 14



	GetColumnToJSON:
	----------------
	csv2json.GetColumnToJSON("/home/Records.csv", 3)

	This function gets data from column number 3



	GetJSONByExcludingColumns:
	-------------------------

	csv2json.GetJSONByExcludingColumns("/home/Records.csv", 1, 14)

	This function gets data by excluding column 1 and 14


	
	GetSortedJSON:
	-------------
	
	csv2json.GetSortedJSON("/home/Records.csv", "Country", csv2json.DESCENDING, csv2json.STRING)
	
	This function first sorts and then returns the sorted data. 