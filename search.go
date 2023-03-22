package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Options struct {
	Key              string
	CaseInSensitive  bool
	LinesBeforeMatch int
	LinesAfterMatch  int
}

func Search(rd io.Reader, opt Options) []string {
	// duplicate reader
	var buf bytes.Buffer
	dup := io.TeeReader(rd, &buf)

	// find matching lines
	var i int
	var indices []int
	scn := bufio.NewScanner(dup)
	for scn.Scan() {
		line := scn.Text()
		// infoLog.Println(line)
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
		addIndicesAfterMatch(ctxMap, opt.LinesAfterMatch, i)
	}

	// create matching lines indices (before and after)
	mIndices := mergeSlices(ctxMap, indices)
	var res []string

	// convert indices to slice of strings
	i, j := 0, 0
	scn = bufio.NewScanner(&buf)
	for scn.Scan() && j < len(mIndices) {
		if mIndices[j] == -1 {
			i, j = i + 1, j + 1
			res = append(res, "--")
			continue
		}

		if i != mIndices[j] {
			i += 1
			continue
		}

		line := scn.Text()
		res = append(res, line)
		i, j = i + 1, j + 1
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

		if len(indices) == 1 || len(res) == 0 || res[len(res)-1] == idx-1 {
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
