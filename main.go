package main

import (
	"go-study/pkg/subjects"
)

func main() {
	// promptOptions()
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/albums", getAlbums)
	// http.HandleFunc("/fib", getFib)

	// 1
	// resp1, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	// body1, _ := io.ReadAll(resp1.Body)
	// json1 := string(body1)

	// // fmt.Println(json1)
	

	cafe := subjects.MakeRequests()

	subjects.SaveJsonToFile(cafe, "requests")
	
	// fmt.Println("tamanho", cafe)


	// err := http.ListenAndServe(":3333", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }
}