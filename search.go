package main

import (
	"bufio"
	"io"
	"strings"
)

type Options struct {
	Key              string
	CaseInSensitive  bool
	LinesBeforeMatch int
	LinesAfterMatch int
}

func Search(rd io.Reader, opt Options) []string {
	var indices []int
	var lines []string

	i := 0
	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		line := scn.Text()
		lines = append(lines, line)

		if match(line, opt) {
			indices = append(indices, i)
		}

		i += 1
	}

	// setup base context map
	var ctxMap = map[int][]int{}
	for _, i := range indices {
		ctxMap[i] = []int{i}
	}

	if opt.LinesBeforeMatch > 0 {
		addIndicesBeforeMatch(ctxMap, opt.LinesBeforeMatch)
	} else if opt.LinesAfterMatch > 0 {
		addIndicesAfterMatch(ctxMap, opt.LinesAfterMatch, len(lines))
	}

	mIndices := mergeSlices(ctxMap, indices)

	var res []string
	for _, i := range mIndices {
		if i == -1 {
			res = append(res, "--")
		} else {
			res = append(res, lines[i])
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

func addIndicesBeforeMatch(m map[int][]int, count int) {
	for i, a := range m {
		j, n := i-1, count

		for n != 0 {
			if j < 0 {
				break
			}
			
			a = append([]int{j}, a...)
			j, n = j-1, n-1
		}

		m[i] = a
	}
}

func addIndicesAfterMatch(m map[int][]int, count int, linesLen int) {
	for i, a := range m {
		j, n := i+1, count

		for n != 0 {
			if j == linesLen {
				break
			}
			
			a = append(a, j)
			j, n = j+1, n-1
		}

		m[i] = a
	}
}

func mergeSlices(m map[int][]int, indices []int) []int {
	var res []int

	for _, i := range indices {
		a := m[i]

		res = appendIndices(res, a)
	}

	return res
}

func appendIndices(res, indices []int) []int {
	for _, idx := range indices {
		i := indexOf(res, idx)

		if i != -1 {
			continue
		}

		if len(res) == 0 || res[len(res)-1] == idx-1 {
			res = append(res, idx)
		} else {
			res = append(res, -1, idx)
		}
	}

	return res
}

func indexOf(s []int, v int) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}
