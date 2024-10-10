package subjects

import (
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
)

/**
 * MakeRequests make 4 requests to the same endpoint.
 * and return the responses in a [][]byte.
 * it was used:
 * - sync.WaitGroup to wait for all goroutines to finish;
 * - A channel to collect the responses;
 * - Goroutines to make the requests;
**/ 
func MakeRequestsWithWaitGroupAndChannel () [][]byte {
	client := resty.New()
	var wg = sync.WaitGroup{}
	// Canal para coletar as respostas
	responseChannel := make(chan []byte, 4)
	
	wg.Add(1)
	go func () {
		resp2, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts/1")
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

func MakeRequestsWithWaitGroup () [][]byte {
	client := resty.New()
	var wg = sync.WaitGroup{}
	results := make([][]byte, 4)
	
	wg.Add(1)
	go func () {
		resp1, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts/1")
		fmt.Println(1)
		results[0] = resp1.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp2, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(2)
		results[1] = resp2.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp3, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(3)
		results[2] = resp3.Body()
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		resp4, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(4)
		results[3] = resp4.Body()
		wg.Done()
	}()

	fmt.Println("Waiting for all requests to finish...")
	
	// Wait for all goroutines to finish
	wg.Wait()

	return results
}

// func MakeRequestsWithErrGroup () [][]byte {
// 	client := resty.New()
// 	var wg = sync.WaitGroup{}
// 	// Canal para coletar as respostas
// 	responseChannel := make(chan []byte, 4)
	
// 	wg.Add(1)
// 	go func () {
// 		resp2, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts/1")
// 		fmt.Println(1)
// 		responseChannel <- resp2.Body()
// 		wg.Done()
// 	}()

// 	wg.Add(1)
// 	go func () {
// 		resp3, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
// 		fmt.Println(2)
// 		responseChannel <- resp3.Body()
// 		wg.Done()
// 	}()

// 	wg.Add(1)
// 	go func () {
// 		resp4, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
// 		fmt.Println(3)
// 		responseChannel <- resp4.Body()
// 		wg.Done()
// 	}()

// 	wg.Add(1)
// 	go func () {
// 		resp5, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts")
// 		fmt.Println(4)
// 		responseChannel <- resp5.Body()
// 		wg.Done()
// 	}()

// 	fmt.Println("Waiting for all requests to finish...")
	
// 	// Espera todas as goroutines terminarem
// 	wg.Wait()
// 	// Fecha o canal, pois todas as goroutines terminaram
// 	close(responseChannel)


// 	// Coletamos todas as respostas
// 	var allResponses [][]byte
// 	for response := range responseChannel {
// 		allResponses = append(allResponses, response)
// 	}

// 	return allResponses
// }