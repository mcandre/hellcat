package main

import (
	"github.com/mcandre/hellcat"

	"flag"
	"fmt"
	"os"
)

var flagExamine = flag.Bool("x", false, "Force hexadecimal dump for text files")
var flagRecurse = flag.Bool("r", false, "Recurse over directories")
var flagHelp = flag.Bool("h", false, "Show usage information")
var flagVersion = flag.Bool("v", false, "Show version information")

func main() {
	flag.Parse()

	var examine bool
	var recurse bool

	if *flagExamine {
		examine = true
	}

	if *flagRecurse {
		recurse = true
	}

	switch {
	case *flagVersion:
		fmt.Println(hellcat.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	toplevels := flag.Args()

	if len(toplevels) == 0 {
		toplevels = []string{"."}
	}

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	config := hellcat.Config{
		Working:   cwd,
		Toplevels: toplevels,
		Examine:   examine,
		Recurse:   recurse,
	}

	if err := config.Roam(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			fmt.Fprintf(os.Stderr, "Error loading path: %v\n", e.Path)
			os.Exit(1)
		}

		panic(err)
	}
}
