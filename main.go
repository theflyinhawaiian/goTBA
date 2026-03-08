package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	inputs := make(chan PlayerChoice)
	events := Start(inputs)

	var input string

	for event := range events {
		fmt.Printf("In state %s\n\n", event.Type.String())
		fmt.Println(event.Description)
		for i, option := range event.Choices {
			fmt.Printf("%d. %s\n", i+1, option.DisplayText)
		}

		fmt.Print("What do? ")
		fmt.Scanln(&input)

		choiceIdx := -1
	GetChoice:
		for i, option := range event.Choices {
			for _, str := range option.Aliases {
				if input == str {
					choiceIdx = i
					break GetChoice
				}
			}
		}

		if choiceIdx == -1 {
			inputValue, err := strconv.Atoi(input)

			if err == nil {
				choiceIdx = inputValue - 1
			}
		}

		if choiceIdx != -1 {
			inputs <- event.Choices[choiceIdx]
		} else {
			fmt.Println("try again!")
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
