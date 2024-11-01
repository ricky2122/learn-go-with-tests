package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Chris")
	_ = http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
