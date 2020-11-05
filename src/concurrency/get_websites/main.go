package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// var wg sync.WaitGroup

func main() {

	start := time.Now()

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	var wg sync.WaitGroup

	for _, website := range websites {
		go getWebsiteV2(website, &wg)
		wg.Add(1)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)

}

func getWebsite(website string) {

	if res, err := http.Get(website); err == nil {
		fmt.Printf("[%d] - %s is up\n", res.StatusCode, website)
	} else {
		fmt.Printf("[%d] - %s is down\n", res.StatusCode, website)
	}

}
func getWebsiteV2(website string, wg *sync.WaitGroup) {
	defer wg.Done()
	if res, err := http.Get(website); err == nil {
		fmt.Printf("[%d] - %s is up\n", res.StatusCode, website)
	} else {
		fmt.Printf("[%d] - %s is down\n", res.StatusCode, website)
	}

}
