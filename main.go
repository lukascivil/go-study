package main

import (
	"go-study/pkg/subjects"
)

func main() {
	// promptOptions()
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/albums", getAlbums)
	// http.HandleFunc("/fib", getFib)
	
	responses := subjects.MakeRequestsWithErrorGroup()

	subjects.SaveJsonToFile(responses, "requests")
}