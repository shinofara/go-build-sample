package main

import (
	"fmt"
  "os"
	"flag"
)

func main() {
	fmt.Println(os.Args)

	// (name, default, help)
	var f = flag.Int("f", 1234, "help message for f")
	var s = flag.Bool("s", false, "help message for bool")
	var long = flag.String("long-option", "default", "help message for long")
	flag.Parse()

	fmt.Println(*f)
	fmt.Println(*s)
	fmt.Println(*long)
}
