package subjects

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJsonToFile(listOfJson [][]byte, fileName string) {
	var jsonObjects []interface{}

	for _, v := range listOfJson {
		var array []interface{}
		var object map[string]interface{}

		err := json.Unmarshal(v, &array)

		if(err != nil) {
			json.Unmarshal(v, &object)

			jsonObjects = append(jsonObjects, object)
		} else {
			jsonObjects = append(jsonObjects, array)
		}
	}

	finalFileName := fmt.Sprintf("%s.json", fileName)
	jsonData, _ := json.MarshalIndent(jsonObjects, "", "  ")

	os.WriteFile(finalFileName, jsonData, 0644)
}