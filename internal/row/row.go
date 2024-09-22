package row

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type die struct {
	sideCount int
}

func (d *die)Roll() int {
	return 1 + rand.Intn(d.sideCount)
}

type row struct {
	dice []die
}

func (r *row)Roll() int {
	sum := 0
	for _, die := range r.dice {
		sum += die.Roll()
	}
	return sum
}

func RollRows(rows *[]row) string {
	res := ""
	for _, row := range *rows {
		res += strconv.Itoa(row.Roll())
		res += "\n"
	}
	return res
}

func ParseString(s string) ([]row, error) {
	var rows []row
	for _, line := range strings.Split(s, "\n") {
		var row row
		for _, diceGroup := range strings.Fields(line) {
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
		rows = append(rows, row)
	}
	return rows, nil
}
