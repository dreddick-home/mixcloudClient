package mixcloud

import (
	"testing"

	"github.com/matryer/is"
	"gotest.tools/assert"
)

var results = Result{
	Data: []Item{
		{
			Key: "this-is-not-wanted",
		},
		{
			Key: "this-is-ok",
		},
		{
			Key: "you-must-have-this",
		},
		{
			Key: "i-have-this-but-also-not-wanted",
		},
	},
}

var filtertests = []struct {
	in  *Filter
	out struct {
		length int
		match  string
	}
}{
	{
		NewFilter(
			[]string{"not wanted"},
			nil,
		),
		struct {
			length int
			match  string
		}{2, ""},
	}, {
		NewFilter(
			nil,
			[]string{"have this"},
		),
		struct {
			length int
			match  string
		}{2, ""},
	}, {
		NewFilter(
			[]string{"not wanted"},
			[]string{"have this"},
		),
		struct {
			length int
			match  string
		}{1, "you-must-have-this"},
	},
}

// Test the offseturl function gives us a valid search string
func TestOffsetURL(t *testing.T) {
	is := is.New(t)
	is.Equal(NewClient("socks", &Filter{}).GetOffsetURL(100), "https://api.mixcloud.com//search/?q=socks&limit=100&offset=100&type=cloudcast")

}

// Test the number of results we get back is as expected
func TestAsyncSearchByCount(t *testing.T) {
	is := is.New(t)
	client := NewClient("babicz", nil)
	client.SearchAsync(0, 5)
	is.Equal(len(client.GetMixItems()), 100)
}

// Benchmark of a simple search
func BenchmarkAsyncSearch(b *testing.B) {
	client := NewClient("sasha", nil)
	client.SearchAsync(15, 5)
}

/* func TestFilterResults(t *testing.T) {
	f := NewFilter(
		[]string{"not wanted"},
		[]string{},
	)
	client := NewClient("foo", f)
	results = client.FilterResults(results)
	assert.Equal(t, len(results.Data), 2)
	assert.Equal(t, results.Data[0].Key, "this-is-ok")
} */

func TestFilterResultsTable(t *testing.T) {
	for i, tt := range filtertests {
		t.Run(string(i), func(t *testing.T) {
			client := NewClient("foo", tt.in)
			r := client.FilterResults(results)
			assert.Equal(t, len(r.Data), tt.out.length)
		})
	}
}
