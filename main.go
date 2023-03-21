package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infoLog = log.New(os.Stdout, "", 0)
	errLog  = log.New(os.Stderr, "", 0)
)

func main() {
	var options struct {
		i       bool
		A, B, C int
		o       string
	}

	flag.BoolVar(&options.i, "i", false, "case insensitize match")
	flag.IntVar(&options.A, "A", 0, "print NUM lines of before match")
	flag.IntVar(&options.B, "B", 0, "print NUM lines of after match")
	flag.IntVar(&options.C, "C", 0, "print count of matches")
	flag.StringVar(&options.o, "o", "", "file to write output")
	flag.Parse()

	// setup output destination
	var output = setOuputDst(options.o)

	if len(flag.Args()) == 0 {
		infoLog.Println("grep: search key arg required")
	}

	// input from stdin
	if len(flag.Args()) == 1 {
		res := Search(os.Stdin, Options{
			Key:              flag.Arg(0),
			LinesAfterMatch:  options.A,
			LinesBeforeMatch: options.B,
			CaseInSensitive:  options.i,
		})

		for _, l := range res {
			fmt.Fprintln(output, l)
		}

		return
	}
}

func setOuputDst(fn string) io.Writer {
	if fn != "" {
		f, err := os.Create(fn)
		if err != nil {
			errLog.Fatalf("grep: %v", err)
		}

		return f
	}
	return os.Stdout
}
