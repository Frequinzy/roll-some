package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	diceSpec := os.Args[1]
	rows, err := ParseString(diceSpec)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row.Roll())
	}
}

type Die struct {
	sideCount int
}

func (d *Die)Roll() int {
	return 1 + rand.Intn(d.sideCount)
}

type Row struct { 
	dice []Die
}

func (r *Row)Roll() int {
	sum := 0
	for _, die := range r.dice {
		sum += die.Roll()
	}
	return sum
}

func ParseString(s string) ([]Row, error) {
	var rows []Row
	for _, line := range strings.Split(s, "\n") {
		var row Row
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
				row.dice = append(row.dice, Die{ sideCount: sideCount})
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}
