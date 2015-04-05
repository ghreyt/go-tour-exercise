package main

import (
	"fmt"
)

// Fetcher
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type FetchResult struct {
	urls  []string
	depth int
}

func Fetch(url string, depth int, resultCh chan FetchResult) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		resultCh <- FetchResult{nil, depth + 1}
		return
	}
	fmt.Println("fetched", depth, url, body)
	resultCh <- FetchResult{urls, depth + 1}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	fetched := make(map[string]bool)
	resultCh := make(chan FetchResult)

	// fetch the root page
	go Fetch(url, 0, resultCh)
	fetching := 1
	fetched[url] = true

	for fetching > 0 {
		result := <-resultCh
		fetching--

		if result.depth > depth {
			// skip deep urls
			continue
		}

		for _, url := range result.urls {
			if !fetched[url] {
				go Fetch(url, result.depth, resultCh)
				fetched[url] = true
				fetching++
			}
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher
// just return already defined body and urls
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
