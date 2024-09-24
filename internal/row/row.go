package row

import (
	"math/rand"
	"regexp"
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
