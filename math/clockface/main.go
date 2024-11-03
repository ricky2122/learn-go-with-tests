package main

import (
	"os"
	"time"

	"learn-go-with-tests/math"
)

func main() {
	t := time.Now()
	math.SVGWriter(os.Stdout, t)
}
