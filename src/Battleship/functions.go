package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func readLine(prompt string) string{
	fmt.Println(prompt)
	fmt.Print("->")

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
