package subjects

import (
	"encoding/json"
	"fmt"
	"go-study/pkg/types"
	"io"
	"net/http"
	"strconv"
)

func GetFib(w http.ResponseWriter, r *http.Request) {
	param1, err := strconv.Atoi(r.URL.Query().Get("param1"))
	result := fib(param1)

	if err != nil {
		// panic(errors.New("number is too large and has wrapped"))
		http.Error(w, "my own error message", http.StatusNotAcceptable)

		return
	}

	io.WriteString(w, strconv.Itoa(result))
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "hello world\n")
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	data := []types.Album{}
	dataWithNoPointer := types.NewAlbum("success0", "John Coltrane", 56.99)
	// dataWithNoPointer.addItemToAlbum("1", "1")

	data = append(data, dataWithNoPointer)
	data = append(data, dataWithNoPointer)
	data = append(data, types.NewAlbum("success1", "Gerry Mulligan", 17.99))
	data = append(data, types.NewAlbum("success2", "Sarah Vaughan", 39.99))

	SaveJsonToFile(data[0])

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}