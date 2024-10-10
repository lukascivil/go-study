package subjects

import (
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
)

func MakeRequests () [][]byte {
	client := resty.New()
	var wg = sync.WaitGroup{}
	// Canal para coletar as respostas
	responseChannel := make(chan []byte, 4)
	
	wg.Add(1)
	go func () {
		resp2, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(1)
		responseChannel <- resp2.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp3, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(2)
		responseChannel <- resp3.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp4, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(3)
		responseChannel <- resp4.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp5, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(4)
		responseChannel <- resp5.Body()
		wg.Done()
	}()

	fmt.Println("Waiting for all requests to finish...")
	
	// Espera todas as goroutines terminarem
	wg.Wait()
	// Fecha o canal, pois todas as goroutines terminaram
	close(responseChannel)


	// Coletamos todas as respostas
	var allResponses [][]byte
	for response := range responseChannel {
		allResponses = append(allResponses, response)
	}

	return allResponses
}