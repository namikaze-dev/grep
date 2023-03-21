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
		r, i bool
		A, B, C int
	}

	flag.BoolVar(&options.i, "i", false, "case insensitize match")
	flag.BoolVar(&options.r, "r", false, "grep directory")
	flag.IntVar(&options.A, "A", 0, "print NUM lines of before match")
	flag.IntVar(&options.B, "B", 0, "print NUM lines of after match")
	flag.IntVar(&options.C, "C", 0, "print count of matches")
	flag.Parse()

	infoLog.Println(flag.Args())
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
