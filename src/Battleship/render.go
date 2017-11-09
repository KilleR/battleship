package main

func render(b *GameBoard) string {
	output := "";

	for _, x := range b.Grid {
		for _,y := range x {
			if y != nil {
				output += "╬"
			} else {
				output += "▒"
			}
		}
		output += "\n"
	}

	return output
}
