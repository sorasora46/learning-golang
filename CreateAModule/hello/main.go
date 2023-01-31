package main

/*
To redirect Go tools from its path (where the module is not there)
to the local directory (where it is)

$ go mod edit -replace example/greeting=../greeting
$ go mod tidy
*/

/*
https://go.dev/doc/tutorial/compile-install
$ go run file.go // compile and run the file without generating binary
$ go build // compile the package, and their dependencies but not install the results
$ go install // compile and install packages
$ go list -f '{{.Target}}' discover the install path
*/

import (
	"example/greeting"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	message, err := greeting.Hello("Sorrawit")
	if err != nil {
		log.Fatal(err) // if err is not nil then print the error message
		// and exit the program
	}
	fmt.Println(message)

	names := []string{"Sorrawit", "Boss", "John"}
	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
