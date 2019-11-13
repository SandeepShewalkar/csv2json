# csv2json

This small package will be helpful to covert csv file data to json array.

# Sample

func main() {

	jsonData := csv2json.ConverToJSONArray("/home/Records.csv")

	for _, val := range jsonData {
		fmt.Println(string(val))
	}
}
