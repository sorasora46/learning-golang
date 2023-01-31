package greeting

// Ending a file's name with _test.go tells the go test command that this file contains test functions.

import (
	"regexp"
	"testing"
)

/*
Test function names have the form TestName,
where Name says something about the specific test.
Also, test functions take a pointer to the testing package's testing.T type as a parameter.
You use this parameter's methods for reporting and logging from your test.
*/

/*
TestHelloName calls the Hello function,
passing a name value with which the function should be able to return a valid response message.
If the call returns an error or an unexpected response message (one that doesn't include the name you passed in),
you use the t parameter's Fatalf method to print a message to the console and end execution.
*/

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

/*
This test is designed to confirm that your error handling works.
If the call returns a non-empty string or no error,
you use the t parameter's Fatalf method to print a message to the console and end execution.
*/

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// in the command line run the "go test"

/*
The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go).
You can add the -v flag to get verbose output that lists all of the tests and their results.
*/
