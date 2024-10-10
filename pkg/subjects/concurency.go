package subjects

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-resty/resty/v2"
	"golang.org/x/sync/errgroup"
)

var client = resty.New()

/**
 * MakeRequests make 4 requests to the same endpoint.
 * and return the responses in a [][]byte.
 * It was used:
**/

/**
 * It was used:
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
	
	// Wait for all goroutines to finish
	wg.Wait()
	
	// Close the channel
	close(responseChannel)

	fmt.Println("Finished!")


	// Coletamos todas as respostas
	var allResponses [][]byte
	for response := range responseChannel {
		allResponses = append(allResponses, response)
	}

	return allResponses
}

/**
 * It was used:
 * - sync.WaitGroup to wait for all goroutines to finish;
 * - Goroutines to make the requests;
**/ 
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

	fmt.Println("Finished!")

	return results
}

/**
 * It was used:
 * - errorGroup to wait for all goroutines to finish and handle errors;
 * - Goroutines to make the requests;
**/ 
func MakeRequestsWithErrorGroup () [][]byte {
	results := make([][]byte, 4)
	errorGroup, ctx := errgroup.WithContext(context.Background())

	errorGroup.Go(func () error{
		resp1, _ := client.R().Get("https://jsonplaceholder.typicode.com/posts/1")
		fmt.Println(1)
		results[0] = resp1.Body()

		return nil
	})

	errorGroup.Go(func () error{
		resp2, _ := client.R().SetContext(ctx).Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(2)
		results[1] = resp2.Body()

		return nil
	})

	errorGroup.Go(func () error{
		resp3, _ := client.R().SetContext(ctx).Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(3)
		results[2] = resp3.Body()

		return nil
	})

	errorGroup.Go(func () error{
		resp4, _ := client.R().SetContext(ctx).Get("https://jsonplaceholder.typicode.com/posts")
		fmt.Println(4)
		results[3] = resp4.Body()

		return nil
	})

	fmt.Println("Waiting for all requests to finish...")
	
	// Wait for all goroutines to finish
	if err := errorGroup.Wait(); err != nil {
		log.Fatalf("Error on HTTP request: %v", err)
	}

	fmt.Println("Finished!")

	return results
}