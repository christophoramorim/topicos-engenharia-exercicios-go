package main

import "fmt"

func isLegalAge(age int) {
	if age < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Printl(isLegalAge(18))
}
