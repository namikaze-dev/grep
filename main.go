package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var (
	infoLog = log.New(os.Stdout, "", 0)
	errLog  = log.New(os.Stderr, "", 0)
)

var options struct {
	i, C bool
	A, B int
	o    string
}

func main() {
	flag.BoolVar(&options.i, "i", false, "case insensitize match")
	flag.IntVar(&options.A, "A", 0, "print NUM lines of before match")
	flag.IntVar(&options.B, "B", 0, "print NUM lines of after match")
	flag.BoolVar(&options.C, "C", false, "print count of matches")
	flag.StringVar(&options.o, "o", "", "file to write output")
	flag.Parse()

	// setup output destination
	var output = setOuputDst(options.o)

	if len(flag.Args()) == 0 {
		infoLog.Println("grep: search key arg required")
	}

	// input from stdin
	if len(flag.Args()) == 1 {
		printSearchResult(os.Stdin, output, "", Options{
			Key:              flag.Arg(0),
			LinesAfterMatch:  options.A,
			LinesBeforeMatch: options.B,
			CaseInSensitive:  options.i,
		})
		return
	}

	for _, fn := range flag.Args()[1:] {
		f, err := os.Open(fn)
		if err != nil {
			errLog.Printf("grep: %v\n", err)
			continue
		}

		fs, err := f.Stat()
		if err != nil {
			errLog.Printf("grep: %v\n", err)
			continue
		}

		if fs.IsDir() {
			searchDir(fn, output)
		}

		var prefix string
		if len(flag.Args()) > 2 {
			prefix = fn + ":"
		}

		printSearchResult(f, output, prefix, Options{
			Key:              flag.Arg(0),
			LinesAfterMatch:  options.A,
			LinesBeforeMatch: options.B,
			CaseInSensitive:  options.i,
		})
		f.Close()
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

func printSearchResult(i io.Reader, o io.Writer, prefix string, opt Options) {
	res := Search(i, opt)

	if options.C {
		fmt.Fprintf(o, "%v%v\n", prefix, len(res))
		return
	}

	for _, l := range res {
		fmt.Fprintf(o, "%v%v\n", prefix, l)
	}
}

func searchDir(dir string, o io.Writer) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			errLog.Printf("grep: %v\n", err)
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			errLog.Printf("grep: %v\n", err)
			return nil
		}

		printSearchResult(f, o, path + ":", Options{
			Key:              flag.Arg(0),
			LinesAfterMatch:  options.A,
			LinesBeforeMatch: options.B,
			CaseInSensitive:  options.i,
		})
		f.Close()

		return nil
	})
}