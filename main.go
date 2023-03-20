package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	infoLog = log.New(os.Stdout, "", 0)
	errLog  = log.New(os.Stderr, "", 0)
)

func main() {
	var options struct {
		A, B, C, r, i bool
	}

	flag.BoolVar(&options.i, "i", false, "case insensitize match")
	flag.BoolVar(&options.r, "r", false, "grep directory")
	flag.BoolVar(&options.A, "A", false, "print NUM lines of before match")
	flag.BoolVar(&options.B, "B", false, "print NUM lines of after match")
	flag.BoolVar(&options.C, "C", false, "print count of matches")
	flag.Parse()

	// args := filterFlags(os.Args[1:])
}

func filterFlags(args []string) []string {
	var bp int
	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			bp = i
			break
		}

		if i == len(args)-1 {
			bp = len(args)
		}
	}

	return args[bp:]
}
