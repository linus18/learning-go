package trans

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	merchant = 3
	amount   = 5
)

//File holds data
type File struct {
	lines      [][]string
	totByMerch map[string]int64
}

//Create is a constructor
func Create(lines [][]string, filter Filter) *File {
	m := make(map[string]int64)
	for _, s := range lines {
		if len(s) != 7 {
			panic("bad data")
		}
		if (filter != nil && !filter(s[0])) || len(s[amount]) == 0 {
			continue
		}
		amt, err := strconv.ParseInt(strings.Replace(s[amount], ".", "", 1), 10, 64)
		if err != nil {
			panic("can't parse amount: ")
		}
		m[s[merchant]] += amt
	}
	return &File{lines, m}
}

//NumOfLines returns number of lines
func (f *File) NumOfLines() int {
	return len(f.lines)
}

//Total returns total amount for a merchant
func (f *File) Total(merchant string) int64 {
	return f.totByMerch[merchant]
}

//Filter function returns true for each line to be included in processing
type Filter func(string) bool

//Print pretty-prints totals by merchant
// TODO: modify to print totals by category for all months available in the data set
func (f *File) Print() {
	for k, v := range f.totByMerch {
		fmt.Printf("%s: $%.*f\n", k, 2, float64(v)/100)
	}
}
