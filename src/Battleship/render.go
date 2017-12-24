package main

import (
)

func render(b *GameBoard) string {
	output := "";

	rows := make([]string, 10)

	for _, x := range b.Grid {
		for j,y := range x {
			if y != nil {
				rows[j] += "╬"
			} else {
				rows[j] += "."
			}
		}
		//output += "\n"
	}

	for _, r := range rows {
		output += r+"\n"
	}

	return output
}
