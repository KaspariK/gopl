// ex1.10 fetches URLs asynchronously and outputs response time, number of bytes in body, and url
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, u := range os.Args[1:] {
		go fetch(u, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(u string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	pu, err := url.Parse(u)
	if err != nil {
		ch <- fmt.Sprintf("while parsing %s: %v", u, err)
		return
	}

	f, err := os.Create(pu.Host)
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if closeErr := f.Close(); closeErr != nil {
		err = closeErr
	}
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", u, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, u)
}

