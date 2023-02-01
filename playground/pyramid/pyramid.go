package pyramid

import "fmt"

func Pyramid() {
	var n int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
