package subjects

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJsonToFile(listOfJson [][]byte, fileName string) {
	var jsonObjects []interface{}

	for _, v := range listOfJson {
		var obj []interface{}

		json.Unmarshal(v, &obj)
		jsonObjects = append(jsonObjects, obj)
	}

	finalFileName := fmt.Sprintf("%s.json", fileName)
	jsonData, _ := json.MarshalIndent(jsonObjects, "", "  ")

	os.WriteFile(finalFileName, jsonData, 0644)
}