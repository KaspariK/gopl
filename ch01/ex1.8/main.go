// ex1.8 fetches page content from URLs passed in through the command line and
// displays it in Stdout
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
