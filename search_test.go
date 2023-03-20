package main_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/namikaze-dev/grep"
)

func TestSearch(t *testing.T) {
	content := `
sample text 
Sample TEXT 
foo
Foobar
FOObaz
	`
	rd := strings.NewReader(content)
	options := main.Options{
		Key: "foo",
		CaseInSensitive: true,
	}

	got := main.Search(rd, options)
	want := []string{"foo", "Foobar", "FOObaz"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	rd = strings.NewReader(content)
	options = main.Options{
		Key: "foo",
		CaseInSensitive: false,
	}

	got = main.Search(rd, options)
	want = []string{"foo"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
