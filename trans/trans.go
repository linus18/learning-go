package trans

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	processingDt = 0
	merchant     = 3
	amount       = 5
)

//File holds data
type File struct {
	lines           [][]string
	totByMonthMerch map[string]map[string]int64
}

//Create is a constructor
func Create(lines [][]string, filter Filter) *File {
	m := make(map[string]map[string]int64)
	for _, s := range lines {
		if len(s) != 7 {
			panic(errors.New("bad data"))
		}
		if (filter != nil && !filter(s[0])) || len(s[amount]) == 0 {
			continue
		}
		date := s[processingDt][0:7]
		totByMerchant := m[date]
		if totByMerchant == nil {
			totByMerchant = make(map[string]int64)
			m[date] = totByMerchant
		}
		amt, err := strconv.ParseInt(strings.Replace(s[amount], ".", "", 1), 10, 64)
		if err != nil {
			panic(errors.New("can't parse amount"))
		}
		totByMerchant[s[merchant]] += amt
	}
	return &File{lines, m}
}

//NumOfLines returns number of lines
func (f *File) NumOfLines() int {
	return len(f.lines)
}

//Filter function returns true for each line to be included in processing
type Filter func(string) bool

//Cruncher knows how to crunch numbers
type Cruncher interface {
	Crunch() map[string]map[string]int64
}

//Crunch builds a map of maps, where keys of the first (outer) are dates
//and the keys of the second one are metchant names
func (f *File) Crunch() map[string]map[string]int64 {
	return f.totByMonthMerch
}

//Print pretty-prints totals by month, merchant
func Print(cruncher Cruncher) {
	for k, v := range cruncher.Crunch() {
		fmt.Printf("Displaying breakdown by merchant for %s:\n", k)
		fmt.Println("============================================")
		for k2, v2 := range v {
			fmt.Printf("%s: $%.*f\n", k2, 2, float64(v2)/100)
		}
	}
}
