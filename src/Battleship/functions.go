package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"errors"
)

func readLine(prompt string) string{
	fmt.Print(prompt)
	//fmt.Print("->")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	// strip LF
	text = strings.Replace(text, "\n", "", -1)
	// strip CR
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func stringToCoords(coord string) ([2]int, error) {
	var output [2]int

	coordRex := regexp.MustCompile(`([a-jA-J])([0-9]{1,2})`)

	coordString := coordRex.FindAllStringSubmatch(coord, -1)

	output[1] = 4
	if len(coordString) > 0 {
		var err error
		switch coordString[0][1] {
		case "a","A":
			output[0] = 0
		case "b","B":
			output[0] = 1
		case "c","C":
			output[0] = 2
		case "d","D":
			output[0] = 3
		case "e","E":
			output[0] = 4
		case "f","F":
			output[0] = 5
		case "g","G":
			output[0] = 6
		case "h","H":
			output[0] = 7
		case "i","I":
			output[0] = 8
		case "j","J":
			output[0] = 9
		default:
			return output, errors.New("First coordinate is invalid (A-J): "+coordString[0][1])
		}
		output[1], err = strconv.Atoi(coordString[0][2])
		if output[1] <1 || output[1] > 10 {
			return output, errors.New("second coordinate out of range, must be 1-10")
		}
		output[1]--
		if err != nil {
			return output, err
		}
		return output, nil
	} else {
		return output, errors.New("Invalid coordinate: "+coord)
	}

}

func coordsToString(coord [2]int) (string, error) {
	output := ""

	if len(coord) == 2 {
		var err error
		switch coord[0] {
		case 0:
			output += "A"
		case 1:
			output += "B"
		case 2:
			output += "C"
		case 3:
			output += "D"
		case 4:
			output += "E"
		case 5:
			output += "F"
		case 6:
			output += "G"
		case 7:
			output += "H"
		case 8:
			output += "I"
		case 9:
			output += "J"
		default:
			return output, errors.New("First coordinate is invalid (0-9): "+strconv.Itoa(coord[0]))
		}
		output += strconv.Itoa(coord[1]+1)
		if err != nil {
			return output, err
		}
		return output, nil
	} else {
		errString := "Invalid coordinates: "
		for i,v := range coord {
			if i != 0 {
				errString += ","
			}
			errString+= strconv.Itoa(v)
		}
		return output, errors.New(errString)
	}

}

func makePlayerShips() []*Ship {
	ships := []*Ship{
		{
			Name:   "Carrier",
			Length: 5,
			Hits:   0,
		},
		{
			Name:   "Battleship",
			Length: 4,
			Hits:   0,
		},
		{
			Name:   "Submarine",
			Length: 3,
			Hits:   0,
		},
		{
			Name:   "Cruiser",
			Length: 3,
			Hits:   0,
		},
		{
			Name:   "Destroyer",
			Length: 2,
			Hits:   0,
		},
	}

	return ships
}
