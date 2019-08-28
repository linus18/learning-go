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
	filterByProcessingDt := func(s string) bool {
		return strings.HasPrefix(s, "2019-07")
	}
	a := Create(data, filterByProcessingDt)
	fmt.Printf("num of lines is %d.\n\n", a.NumOfLines())
	tot := a.Crunch()["2019-07"]["Sunny Cafe - Anyplace USA"]
	if tot != 3432 {
		t.Errorf("total amount should be %d but is %d", 3432, tot)
	}
	tot = a.Crunch()["2019-07"]["Larry The Cable Guy"]
	fmt.Println(tot)
	Print(a)
	if m := a.Crunch()["2019-08"]; m != nil {
		t.Errorf("data for 2019-08 should be nil but is %v", m)
	}
	b := Create(data, nil)
	tot = b.Crunch()["2019-08"]["Sunny Cafe - Anyplace USA"]
	if tot != 1321 {
		t.Errorf("total should be %d but is %d", 1321, tot)
	}
	if m := b.Crunch()["2019-07"]; m == nil {
		t.Error("data for 2019-07 should not be nil")
	}
	Print(b)
}

func TestBadData(t *testing.T) {
	assertPanic(t, "bad data", func() {
		Create([][]string{{""}}, nil)
	})
}

func TestBadData2(t *testing.T) {
	assertPanic(t, "can't parse amount", func() {
		Create([][]string{{"2019-07-01", "", "", "", "", "XYZ", ""}}, nil)
	})
}

func assertPanic(t *testing.T, msg string, f func()) {
	defer func() {
		err := recover().(error)
		if err == nil {
			t.Errorf("should've failed")
		}
		if err.Error() != msg {
			t.Errorf("Wrong panic message: %s,", err.Error())
		}
	}()
	f()
}

func TestCrunch(t *testing.T) {
	c := Create(data, nil)
	Print(c)
	data := c.Crunch()
	tot := data["2019-07"]["Sunny Cafe - Anyplace USA"]
	if tot != 3432 {
		t.Errorf("total should be %d but is %d", 3432, tot)
	}
}
