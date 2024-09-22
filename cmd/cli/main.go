package main

import (
	"fmt"
	"os"
	"github.com/Frequinzy/roll-some/internal/row"
)

func main() {
	diceSpec := os.Args[1]
	rows, err := row.ParseString(diceSpec)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", row.RollRows(&rows))
}

