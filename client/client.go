package mixcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	apiBase  = "https://api.mixcloud.com/"
	maxIters = 100
)

// Filter is a client side filter for the
// results
type Filter struct {
	Excludes *regexp.Regexp
	Includes *regexp.Regexp
}

// NewFilter returns a new filter
func NewFilter(excludes []string, includes []string) *Filter {
	var eregex *regexp.Regexp
	if len(excludes) != 0 {
		e := excludes[:0]
		for _, val := range excludes {
			e = append(e, ConvertSearchString(val))
		}
		eregex = regexp.MustCompile(strings.Join(e, "|"))
	}

	var iregex *regexp.Regexp
	if len(includes) != 0 {
		i := includes[:0]
		for _, val := range includes {
			i = append(i, ConvertSearchString(val))
		}
		iregex = regexp.MustCompile(strings.Join(i, "|"))
	}

	return &Filter{
		Excludes: eregex,
		Includes: iregex,
	}

}

// ConvertSearchString converts a search term to match
// those found in mixcloud results
func ConvertSearchString(s string) string {
	var re = regexp.MustCompile(" ")
	term := re.ReplaceAllString(s, "-")
	return term
}

// NewClient instantiates a new Client
// using the search term
func NewClient(term string, filter *Filter) *Client {
	term = ConvertSearchString(term)
	return &Client{
		apiBase: apiBase,
		resultSet: ResultSet{
			[]Result{},
			&sync.Mutex{},
		},
		searchTerm: term,
		filter:     filter,
	}
}

// Client is a client used to interact with the Mixcloud API
type Client struct {
	// unexported fields...
	apiBase    string
	resultSet  ResultSet
	searchTerm string
	filter     *Filter
}

// Item is a search result for a single mix
type Item struct {
	Key string `json:"key"` // Key is the mix key field of the item
}

// Result is the response back from the API
type Result struct {
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
	Data []Item `json:"data"`
}

// ResultSet holds a slice of Results including a mutex
// for thread safe updates.
type ResultSet struct {
	results []Result
	mutex   *sync.Mutex
}

// PrintResults prints results to stdout
// should be replaced by a func that returns string
func (c *Client) PrintResults() {
	for _, val := range c.resultSet.results {
		for _, nal := range val.Data {
			fmt.Println(nal.Key)
		}
	}
}

// GetMixItems returns a slice of strings which are the keys of the mix data
func (c *Client) GetMixItems() []string {
	s := []string{}
	for _, val := range c.resultSet.results {
		for _, nal := range val.Data {
			s = append(s, nal.Key)
		}
	}
	return s

}

func (r *ResultSet) update(result Result) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.results = append(r.results, result)
	return nil
}

// GetOffsetURL returns a mixcloud search URL from the offset
// passed in.
func (c *Client) GetOffsetURL(offset int32) string {
	return fmt.Sprintf("%s/search/?q=%s&limit=100&offset=%d&type=cloudcast", apiBase, c.searchTerm, offset)
}

type queueItem struct {
	integer int32
	done    bool
}

func (c *Client) loadQueue(max int32, wg *sync.WaitGroup, q chan queueItem) {
	for i := int32(0); i < max; i++ {
		q <- queueItem{integer: i * 100, done: false}
	}
	q <- queueItem{done: true}
	wg.Done()
}

// SearchAsync runs up to 5 workers asynchronously
// to search mixcloud using a search term and a range of offsets
func (c *Client) SearchAsync(max int32, workers int32) {

	// We put a hard limit on the number of iterations to prevent too many
	if max > maxIters {
		max = maxIters
	}
	var wg sync.WaitGroup
	queueChan := make(chan queueItem, workers)

	wg.Add(1)
	go c.loadQueue(max, &wg, queueChan)

	c.GetAsync(&wg, queueChan)

	wg.Wait()

}

// GetAsync asynchronously GETs a mixcloud URL pulling in items
// from a channel
func (c *Client) GetAsync(wg *sync.WaitGroup, q chan queueItem) {
	for done := false; done == false; {
		item := <-q
		if item.done {
			done = true
		}
		offset := item.integer
		url := c.GetOffsetURL(offset)
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			dec := json.NewDecoder(resp.Body)
			a := Result{}
			dec.Decode(&a)
			if c.filter != nil {
				a = c.FilterResults(a)
			}
			c.resultSet.update(a)
			wg.Done()
		}(url)

	}
}

// FilterResults filters the results according to the filters
// set on the client
func (c *Client) FilterResults(r Result) Result {
	if c.filter == nil {
		return r
	}
	var items = make([]Item, 0, len(r.Data))

	if c.filter.Excludes != nil {
		for _, val := range r.Data {
			if c.filter.Excludes.Match([]byte(val.Key)) == false {
				items = append(items, val)
			}
		}
		r.Data = items
	}

	if c.filter.Includes == nil {
		return r
	}

	var itemsinc = make([]Item, 0, len(r.Data))
	for _, val := range r.Data {
		if c.filter.Includes.Match([]byte(val.Key)) == true {

			itemsinc = append(itemsinc, val)
		}
	}

	r.Data = itemsinc

	return r

}
