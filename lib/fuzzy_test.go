package sman

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"reflect"
	"testing"
)

func TestTopsFromRanks(t *testing.T) {
	tests := []struct {
		name        string
		ranks       fuzzy.Ranks
		wantMatched []string
	}{
		{"empty ranks",
			fuzzy.Ranks{},
			[]string(nil),
		},
		{"single matched",
			fuzzy.Ranks{{"fir", "first", 3, 1}, {"fir", "second", 8, 2}},
			[]string{"first"},
		},
		{"multiple matched",
			fuzzy.Ranks{{"fir", "first", 3, 1}, {"fir", "second", 3, 2}, {"fir", "invalid", 8, 3}},
			[]string{"first", "second"},
		},
	}
	for _, tt := range tests {
		if gotMatched := topsFromRanks(tt.ranks); !reflect.DeepEqual(gotMatched, tt.wantMatched) {
			t.Errorf("%q. topsFromRanks() = %v, want %v", tt.name, gotMatched, tt.wantMatched)
		}
	}
}

func TestFSearchSnippet(t *testing.T) {
	tests := []struct {
		name        string
		snippets    SnippetSlice
		pattern     string
		wantMatched SnippetSlice
	}{
		{"no match",
			SnippetSlice{
				Snippet{Name: "first"},
				Snippet{Name: "nonfirst"},
			}, "bird",
			SnippetSlice(nil),
		},
		{"single matched",
			SnippetSlice{
				Snippet{Name: "first"},
				Snippet{Name: "nonfirst"},
			}, "first",
			SnippetSlice{Snippet{Name: "first"}},
		},
		{"multiple matched",
			SnippetSlice{
				Snippet{Name: "first"},
				Snippet{Name: "firbe"},
				Snippet{Name: "non:match"},
			}, "fir",
			SnippetSlice{
				Snippet{Name: "first"},
				Snippet{Name: "firbe"},
			},
		},
	}
	for _, tt := range tests {
		if gotMatched := fSearchSnippet(tt.snippets, tt.pattern); !reflect.DeepEqual(gotMatched, tt.wantMatched) {
			t.Errorf("%q. fSearchSnippet() = %#v, want %#v", tt.name, gotMatched, tt.wantMatched)
		}
	}
}
