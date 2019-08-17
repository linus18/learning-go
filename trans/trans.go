package trans

import (
	"fmt"
	"log"
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

func Create(lines [][]string) *File {
	m := make(map[string]int64)
	for _, s := range lines {
		if len(s[AMOUNT]) == 0 {
			continue
		}
		amt, err := strconv.ParseInt(strings.Replace(s[AMOUNT], ".", "", 1), 10, 64)
		if err != nil {
			log.Fatal("can't parse amount: ", err)
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

func (f *File) Print() {
	for k, v := range f.totByMerch {
		fmt.Printf("%s: $%.*f\n", k, 2, float64(v)/100)
	}
}

//TODO: implement feature to total using filters (month, year, etc.)
func (f *File) TotalWithFilter(merchant string, filter interface{}) int64 {
	return 0
}
