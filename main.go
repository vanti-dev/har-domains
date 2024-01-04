package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: har-domains <har-file>")
	}
	har, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	var data map[string]any
	if err := json.Unmarshal(har, &data); err != nil {
		panic(err)
	}
	urlC := make(chan string, 10)
	go func() {
		findURLs(data, urlC)
		close(urlC)
	}()

	var domains []string // sorted, unique domain names
	for urlStr := range urlC {
		u, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		h := u.Hostname()
		i := sort.SearchStrings(domains, u.Hostname())
		if i < len(domains) && domains[i] == h {
			continue
		}
		domains = append(domains, "")
		copy(domains[i+1:], domains[i:])
		domains[i] = h
	}

	for _, d := range domains {
		println(d)
	}
}

func findURLs(m map[string]any, urlC chan<- string) {
	for k, v := range m {
		switch v := v.(type) {
		case string:
			if k == "url" {
				urlC <- v
			}
		case map[string]any:
			findURLs(v, urlC)
		case []any:
			for _, v := range v {
				if v, ok := v.(map[string]any); ok {
					findURLs(v, urlC)
				}
			}
		}
	}
}
