package main

import (
	"errors"
	"fmt"
)

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(x, y string) {
		fmt.Println(formatter(x, y))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("==========================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbol found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}
