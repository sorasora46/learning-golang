package user

import "fmt"

// if you want to use attribute outside the package
// the attribute should be uppercase too
type Student struct {
	Name  string
	Year  uint
	Major string
}

func (s Student) Greeting() {
	fmt.Printf("Hi! My name is %s\n", s.Name)
}
