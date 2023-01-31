package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Go code is grouped into packages,
// and packages are grouped into modules.

// Your module specifies dependencies needed to run your code,
// including the Go version and the set of other modules it requires.

/* TLDR; go.mod == package.json
The go mod init command creates a go.mod file to track your code's dependencies.
So far, the file includes only the name of your module and the Go version your code supports.
But as you add dependencies, the go.mod file will list the versions your code depends on.
This keeps builds reproducible and gives you direct control over which module versions to use.
*/

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with the messages
	messages := make(map[string]string)
	// create a map in go with below syntax:
	// map[key-type]value-type
	// messages = map with string as key and string as value

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// for loop with "range" return 2 values
	// 1st: index, 2nd copy of item value
	// use "_" to ignore the 1st value
	for _, name := range names {

		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// Go can return multiple values
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// := means let the go compiler determine the type
	message := fmt.Sprintf(randomFormat(), name)
	// Sprintf = return the formatted string
	return message, nil // nil means no error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// slice is dynamic array
	// A slice of message format ([]string : slice format)
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
