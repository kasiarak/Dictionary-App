package main

import (
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentRequests(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://localhost:" + port + "/dictionary")
			assert.NoError(t, err, "Server request failed")
			assert.Equal(t, 200, resp.StatusCode, "Server should return 200 OK")
		}()
	}

	wg.Wait()
}
