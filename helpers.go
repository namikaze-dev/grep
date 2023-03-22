package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

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

func createSearchResult(i io.Reader, prefix string, opt Options) string {
	r := Search(i, opt)

	if options.C {
		return fmt.Sprintf("%v%v\n", prefix, len(r))
	}

	var res []string
	for _, l := range r {
		res = append(res, fmt.Sprintf("%v%v", prefix, l))
	}

	return strings.Join(res, "\n")
}
