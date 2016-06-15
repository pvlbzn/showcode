// showcode -c path/to/source/file
// showcode -d path/to/image
package main

import (
	"./coder"
	"flag"
)

var (
	code   = flag.Bool("c", false, "code a source file")
	decode = flag.Bool("d", false, "decode an image")
)

func main() {
	flag.Parse()

	if *code {
		coder.Code(flag.Args()[0])
	}

	if *decode {
		coder.Decode(flag.Args()[0])
	}
}
