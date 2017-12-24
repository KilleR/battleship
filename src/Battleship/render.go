package main

import (
	"strconv"
)

func render(p *Player) string {
	b := p.Board
	output := "";

	rows := make([]string, 11)

	// put in grid headers
	rows[0] = ".  A B C D E F G H I J"
	for i:=1;i<len(rows);i++ {
		if i<10 {
			rows[i] += " "
		}
		rows[i] += strconv.Itoa(i)
	}

	// put the pieces into rows
	for i, x := range b.Grid {
		for j,y := range x {
			// check for hit tiles
			if p.ShotsReceived[i][j] != nil {
				if *p.ShotsReceived[i][j] == true {
					rows[j+1] += " X"
				} else {
					rows[j+1] += " o"
				}
			} else {
				if y == nil {
					rows[j+1] += " ."
				} else {
					rows[j+1] += " ╬"
				}
			}
		}
	}

	// render the rows
	for _, r := range rows {
		output += r+"\n"
	}

	return output
}

func renderFired(p *Player) string {
	output := "";

	rows := make([]string, 11)

	// put in grid headers
	rows[0] = ".  A B C D E F G H I J"
	for i:=1;i<len(rows);i++ {
		if i<10 {
			rows[i] += " "
		}
		rows[i] += strconv.Itoa(i)
	}

	// put the pieces into rows
	for _, x := range p.ShotsFired {
		for j,y := range x {
			if y == nil {
				rows[j+1] += " ."
			} else {
				if *y == true {
					rows[j+1] += " X"
				} else {
					rows[j+1] += " o"
				}
			}
		}
	}

	// render the rows
	for _, r := range rows {
		output += r+"\n"
	}

	return output
}