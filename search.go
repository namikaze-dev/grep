package main

import (
	"bufio"
	"io"
	"strings"
)

type Options struct {
	Key             string
	CaseInSensitive bool
}

func Search(rd io.Reader, opt Options) []string {
	var res []string

	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		line := scn.Text()

		if match(line, opt) {
			res = append(res, line)
		}
	}

	return res
}

func match(line string, opt Options) bool {
	if opt.CaseInSensitive {
		line = strings.ToLower(line)
		return strings.Contains(line, strings.ToLower(opt.Key))
	} else {
		return strings.Contains(line, opt.Key)
	}
}
