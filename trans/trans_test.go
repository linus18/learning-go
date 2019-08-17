package trans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

var data = make([][]string, 0)

func TestMain(m *testing.M) {
	f, err := os.Open("trans.csv")
	if err != nil {
		panic(err)
	}
	in := bufio.NewScanner(f)
	header := true
	for in.Scan() {
		line := in.Text()
		if header {
			header = false
			continue
		}
		split := strings.Split(line, ",")
		ln := make([]string, len(split))
		for i, s := range split {
			ln[i] = s
		}
		data = append(data, ln)
	}
	defer f.Close()
	os.Exit(m.Run())
}

func TestTotal(t *testing.T) {
	a := Create(data)
	fmt.Printf("num of lines is %d.\n\n", a.NumOfLines())
	tot := a.Total("Sunny Cafe - Anyplace USA")
	if tot != 3432 {
		t.Errorf("total amount should be %d but is %d", 3432, tot)
	}
	tot = a.Total("Larry The Cable Guy")
	fmt.Println(tot)
	a.Print()
}
