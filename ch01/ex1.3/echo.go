// ex1.3 benchmarks various methods for printing command line arguments
package echo

import (
	"strings"
)

func main() {

}

// Loop uses a standard loop
func Loop(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

// Range uses a range loop
func Range(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
}

// Join uses strings.Join
func Join(args []string) {
	strings.Join(args[1:], " ")
}
