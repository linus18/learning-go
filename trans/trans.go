package trans

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MERCHANT = 3
	AMOUNT   = 5
)

type File struct {
	lines      [][]string
	totByMerch map[string]int64
}

func Create(lines [][]string, filter Filter) *File {
	m := make(map[string]int64)
	for _, s := range lines {
		if len(s) != 7 {
			panic("bad data")
		}
		if (filter != nil && !filter(s[0])) || len(s[AMOUNT]) == 0 {
			continue
		}
		amt, err := strconv.ParseInt(strings.Replace(s[AMOUNT], ".", "", 1), 10, 64)
		if err != nil {
			panic("can't parse amount: ")
		}
		m[s[MERCHANT]] += amt
	}
	return &File{lines, m}
}

func (f *File) NumOfLines() int {
	return len(f.lines)
}

func (f *File) Total(merchant string) int64 {
	return f.totByMerch[merchant]
}

type Filter func(string) bool

//TODO: modify to print totals by category for all months available in the data set
func (f *File) Print() {
	for k, v := range f.totByMerch {
		fmt.Printf("%s: $%.*f\n", k, 2, float64(v)/100)
	}
}
