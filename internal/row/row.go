package row

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type die struct {
	sideCount int
	value int
}

func (d *die)Roll() {
	d.value = 1 + rand.Intn(d.sideCount)
}

type row struct {
	dice []die
	modifier int
}


// row = expr | expr, row | expr, operator, row ;
// expr = modifier, action | action | int ;
// action = "d", int | "(", row, ")" ;
// modifier = low | high | int ;
// low = "l", int | int, "l", int ;
// high = "h", int | int, "h", int ;
// operator = "+", | "-" ;
type eRow struct {
	exprs []expr
	operators []operator
}

type expr struct {
	action action
	modifier modifier
}

type num int 

type action interface {
	calculate() int
}

func (r *eRow)calculate() int {
	sum := 0
	for i := range r.exprs {
		e := r.exprs[i]
		sum += e.modifier.modify(e.action)
	}
	fmt.Println(sum)
	return sum
}

func (d *die)calculate() int {
	d.Roll()
	fmt.Println(d.value)
	return d.value
}

func (n *num)calculate() int {
	return int(*n)
}

type modifier interface {
	modify(a action) int
}

type noMod struct {}

func (n *noMod)modify(a action) int {
	return a.calculate()
}

type multiplier int 

func (m multiplier)modify(a action) int {
	sum := 0
	for range m {
		sum += a.calculate()
	}
	return sum
}

type lowest struct {
	pickN int
	ofN int
}

func (l *lowest)modify(a action) int {
	var results sort.IntSlice
	for range l.ofN {
		res := a.calculate()
		results = append(results, res)
	}

	results.Sort()

	sum := 0
	for i := 0; i < l.pickN; {
		sum += results[i]
	}
	return sum
}

type highest struct {
	pickN int
	ofN int
}

func (l *highest)modify(a action) int {
	var results sort.IntSlice
	for range l.ofN {
		res := a.calculate()
		results = append(results, res)
	}

	sort.Sort(sort.Reverse(results))

	sum := 0
	for i := range l.pickN {
		sum += results[i]
	}
	return sum
}

type operator interface {
	operate() int
}

func ParseRow(s string) (*row, error) {
	var row row
	for _, diceGroup := range strings.Fields(s) {
		re := regexp.MustCompile(`(\d*)d(\d+)`)
		matches := re.FindStringSubmatch(diceGroup)
		dieCount, err := strconv.Atoi(matches[1])
		if err != nil {
			dieCount = 1
		}
		sideCount, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}

		for range dieCount {
			row.dice = append(row.dice, die{ sideCount: sideCount})
		}
	}
	return &row, nil
}

func (r *row)SumRow() int {
	sum := r.modifier
	for _, die := range r.dice {
		sum += die.value
	}
	return sum
}

func (r *row)Roll() {
	for i := range r.dice {
		r.dice[i].Roll()
	}
}

func RollRows(rows *[]row) string {
	res := ""
	for _, row := range *rows {
		row.Roll()
		res += strconv.Itoa(row.SumRow())
		res += "\n"
	}
	return res
}

func ParseString(s string) ([]row, error) {
	var rows []row
	for _, line := range strings.Split(s, "\n") {
		row, err := ParseRow(line)
		if err != nil {
			return nil, err
		}
		rows = append(rows, *row)
	}
	return rows, nil
}
