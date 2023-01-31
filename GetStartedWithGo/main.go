package main // a package is a way to group a function
// it's made up of all files in the same directory

import (
	"fmt" // format package

	"rsc.io/quote" // use "go mod tidy" to install missing packages in go.mod and update/create go.sum
)

// Go code is grouped into packages,
// and packages are grouped into modules.
// Your module specifies dependencies needed to run your code,
// including the Go version and the set of other modules it requires.

func main() {
	fmt.Println(quote.Go())
}
