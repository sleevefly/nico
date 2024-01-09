package service

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"sync"
	"testing"
	"time"
)

func makeRequest(client *resty.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := client.R().Get("http://localhost:8080/ping")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", resp.Status())
}

func TestHttpRequest(t *testing.T) {
	client := resty.New()

	var wg sync.WaitGroup

	// Define the number of concurrent requests
	numRequests := 20

	// Make multiple requests concurrently
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go makeRequest(client, &wg)
	}

	// Wait for all requests to finish
	wg.Wait()

	// Allow some time to observe the responses before exiting
	time.Sleep(2 * time.Second)
}
