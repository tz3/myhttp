package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xmoutaz/myhttp/http"
	"github.com/xmoutaz/myhttp/utils"
)

func main() {
	// Parse command-line arguments
	parallel := flag.Int("parallel", 10, "number of parallel requests to perform")
	flag.Parse()

	if *parallel <= 0 {
		fmt.Println("number of parallel should be > 0")
		os.Exit(1)
	}
	// Validate URLs
	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("Usage: myhttp [-parallel n] url1 [url2 [url3 ...]]")
		os.Exit(1)
	}
	for i, url := range urls {
		urls[i] = utils.FixURLScheme(url)
		if err := utils.ValidateURL(urls[i]); err != nil {
			fmt.Printf("%s is not a valid URL: %s\n", url, err)
			continue
		}
	}

	// Make HTTP requests in parallel
	results := make(chan http.Result, len(urls))
	semaphore := make(chan struct{}, *parallel)
	for _, url := range urls {
		semaphore <- struct{}{}
		go func(url string) {
			defer func() { <-semaphore }()
			resp, err := http.Fetch(url)
			if err != nil {
				results <- http.Result{URL: url, Error: err}
				return
			}
			results <- http.Result{URL: url, MD5Hash: utils.MD5Hash(resp)}
		}(url)
	}

	// Print results as they become available
	for i := 0; i < len(urls); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("%s: %s\n", result.URL, result.Error)
		} else {
			fmt.Printf("%s %s\n", result.URL, result.MD5Hash)
		}
	}
}
