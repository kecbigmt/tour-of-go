package main
import (
	"fmt"
  "sync"
)

var wg sync.WaitGroup

func MakeIntersection(set1 []string, set2 []string) []string{
  dict := map[string]bool{}
  for _, v := range set1{
    dict[v] = true
  }
  for _, v := range set2{
    if dict[v] {
      delete(dict, v)
    }
  }
  result := make([]string, len(dict), len(dict))
  i := 0
  for k, _ := range dict {
    result[i] = k
    i++
  }
  return result
}

type History struct{
  urls []string
  mux sync.Mutex
}
func (h *History) Add(u string) {
  h.mux.Lock()
  h.urls = append(h.urls, u)
  h.mux.Unlock()
  return
}
func (h *History) Filter(us []string) []string {
  h.mux.Lock()
  defer h.mux.Unlock()
  return MakeIntersection(us, h.urls)
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, h *History) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
  defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
  urls = h.Filter(urls)
	for _, u := range urls {
    h.Add(u)
    wg.Add(1)
		go Crawl(u, depth-1, fetcher, h)
	}
	return
}

func main() {
  start := "http://golang.org/"
  h := History{urls:[]string{start}}
  wg.Add(1)
	go Crawl(start, 4, fetcher, &h)
  wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
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
