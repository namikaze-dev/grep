package main_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/namikaze-dev/grep"
)

func TestSearchBasic(t *testing.T) {
	content := "sample text\nSample TEXT\nfoo\nFoobar\nFOObaz"
	rd := strings.NewReader(content)
	options := main.Options{
		Key:             "foo",
		CaseInSensitive: true,
	}

	got := main.Search(rd, options)
	want := []string{"foo", "Foobar", "FOObaz"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	rd = strings.NewReader(content)
	options = main.Options{
		Key:             "foo",
		CaseInSensitive: false,
	}

	got = main.Search(rd, options)
	want = []string{"foo"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSearch_LinesBeforeMatch(t *testing.T) {
	t.Run("search with 2 lines before match and case insensitive", func(t *testing.T) {
		content := "sample text\nSample TEXT\nfoo\nFoobar\nFOObaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "foo",
			CaseInSensitive:  true,
			LinesBeforeMatch: 2,
		}

		got := main.Search(rd, options)
		want := []string{"sample text", "Sample TEXT", "foo", "Foobar", "FOObaz"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("search with 2 lines before match and case sensitive", func(t *testing.T) {
		content := "genos\nhelloworld\nhellobar\nfoobar\ngenos\nbaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "genos",
			CaseInSensitive:  false,
			LinesBeforeMatch: 1,
		}

		got := main.Search(rd, options)
		want := []string{"genos", "--", "foobar", "genos"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("search with 2000 lines before match and case sensitive", func(t *testing.T) {
		content := "genos\nhelloworld\nhellobar\nfoobar\ngenos\nbaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "genos",
			CaseInSensitive:  false,
			LinesBeforeMatch: 2000,
		}

		got := main.Search(rd, options)
		want := []string{"genos", "helloworld", "hellobar", "foobar", "genos"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSearch_LinesAfterMatch(t *testing.T) {
	t.Run("search with 2 lines before match and case insensitive", func(t *testing.T) {
		content := "sample text\nSample TEXT\nfoo\nFoobar\nFOObaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "foo",
			CaseInSensitive:  true,
			LinesAfterMatch: 2,
		}

		got := main.Search(rd, options)
		want := []string{"foo", "Foobar", "FOObaz"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("search with 2 lines before match and case sensitive", func(t *testing.T) {
		content := "genos\nhelloworld\nhellobar\nfoobar\ngenos\nbaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "genos",
			CaseInSensitive:  false,
			LinesAfterMatch: 1,
		}

		got := main.Search(rd, options)
		want := []string{"genos", "helloworld", "--", "genos", "baz"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("search with 2000 lines before match and case sensitive", func(t *testing.T) {
		content := "genos\nhelloworld\nhellobar\nfoobar\ngenos\nbaz"
		rd := strings.NewReader(content)
		options := main.Options{
			Key:              "genos",
			CaseInSensitive:  false,
			LinesAfterMatch: 2000,
		}

		got := main.Search(rd, options)
		want := []string{"genos", "helloworld", "hellobar", "foobar", "genos", "baz"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
