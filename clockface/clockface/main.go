package main

import (
	"os"
	"time"

	"github.com/bvkin/go-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
