package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
)

func saveAlbum(album album) {
	// data := []album{}

	albumJson, _ := json.Marshal(album)
	os.WriteFile("albums.json", albumJson, 0644)
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	data := []album{}
	dataWithNoPointer := newAlbum("success0", "John Coltrane", 56.99)
	// dataWithNoPointer.addItemToAlbum("1", "1")

	data = append(data, dataWithNoPointer)
	data = append(data, dataWithNoPointer)
	data = append(data, newAlbum("success1", "Gerry Mulligan", 17.99))
	data = append(data, newAlbum("success2", "Sarah Vaughan", 39.99))

	saveAlbum(data[0])

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func fib (n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func getFib(w http.ResponseWriter, r *http.Request) {
	param1, err := strconv.Atoi(r.URL.Query().Get("param1"))
	result := fib(param1)

	if err != nil {
		// panic(errors.New("number is too large and has wrapped"))
		http.Error(w, "my own error message", http.StatusNotAcceptable)

		return
	}

	io.WriteString(w, strconv.Itoa(result))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "hello world\n")
}

func getInput(prompt string, reader *bufio.Reader) (string) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt := getInput("Choose option (a - add item, s - save album)", reader)

	fmt.Println(opt)
}

func cafe(oi *int) int {
	*oi = 2

	return *oi
}

func makeRequests () []string {
	client := resty.New()
	var wg = sync.WaitGroup{}
	// Canal para coletar as respostas
	responseChannel := make(chan string, 4)
	
	wg.Add(1)
	go func () {
		resp2, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(1)
		responseChannel <- resp2.String()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp3, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(2)
		responseChannel <- resp3.String()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp4, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(3)
		responseChannel <- resp4.String()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp5, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(4)
		responseChannel <- resp5.String()
		wg.Done()
	}()

	fmt.Println("Waiting for all requests to finish...")
	
	// Espera todas as goroutines terminarem
	wg.Wait()
	// Fecha o canal, pois todas as goroutines terminaram
	close(responseChannel)


	// Coletamos todas as respostas
	var allResponses []string
	for response := range responseChannel {
		allResponses = append(allResponses, response)
	}

	return allResponses
}

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
	

	cafe := makeRequests()
	
	fmt.Println("tamanho", cafe)


	// err := http.ListenAndServe(":3333", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }
}