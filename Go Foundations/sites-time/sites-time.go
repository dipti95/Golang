package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	//siteTime("https://google.com")

	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://no-such-site.biz",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		//wg.Add(1)
		url := url

		go func() {
			defer wg.Done()
			siteTime(url)
		}()

	}
	wg.Wait()
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error: %s -> %s", url, err)
		return
	}

	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("error: %s -> %s", url, err)

	}

	duration := time.Since(start)

	log.Printf("Info: %s -> %v", url, duration)
}
