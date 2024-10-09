package subjects

import (
	"encoding/json"
	"os"
)

func SaveJsonToFile[T any](jsonBin T) {
	albumJson, _ := json.Marshal(jsonBin)
	
	os.WriteFile("albums.json", albumJson, 0644)
}