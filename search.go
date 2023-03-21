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
	var indices []int
	var m = map[int]string{}

	i := 0
	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		line := scn.Text()

		if match(line, opt) {
			indices = append(indices, i)
			m[i] = line
		}

		i += 1
	}

	var res []string
	for _, i := range indices {
		res = append(res, m[i])
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
