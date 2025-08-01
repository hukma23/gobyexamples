package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program's execution, usually for purpose of cleanup.
// defer is often used where e.g. ensure and finally would be used in other languages.

func main() {
	// Suppose we wanted to create a file, write to it, and then close when we're done. Here's how we could do that with defer.

	// Immidiately after getting a file object with createFile, we defer the closing of that file with closeFile.
	// This will be executed at the end of the enclosing function (main), after writeFile has finished.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("createing")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// Its important to check for errors when closing a file,even in deferred function.
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}
