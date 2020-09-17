// ex1.2 prints invoked command, arguments, and indexes
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		fmt.Printf("%d: %s\n", i, a)
	}
}
