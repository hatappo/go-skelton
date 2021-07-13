package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Printf("There are '%d' args and those are: %s\n", flag.NArg(), flag.Args())
}
