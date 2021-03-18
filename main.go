package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	AppVersion = "1.0"
	DigitsHeld = 100001
)

var (
	random = flag.Bool("random", false, "Pick a random digit to test your PI memory")
	startingDigit = flag.Int("start", 1, "Starting digit for prompt")
	displayTip = flag.Bool("tipme", false, "Display the digit tip")
	loop = flag.Bool("loop", false, "Keep trying forever until a control signal")
	limitRandom = flag.Int("limit", DigitsHeld, "Limit the digit upper bound for random digits")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("PI Memory Checker version %s\n", AppVersion)
	if *displayTip {
		fmt.Printf("Digits %d - %d is: %s\n", *startingDigit, *startingDigit+5, PI[*startingDigit+1:*startingDigit+6])
	}
	if *random {
		reader := bufio.NewReader(os.Stdin)
		randDigit := rand.Intn(*limitRandom)
		fmt.Printf("Enter digit number [%d]: ", randDigit)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if string(PI[randDigit+1]) == text {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Oops its actually: %s\n", string(PI[randDigit+1]))
		}
		os.Exit(0)
	}

	if *loop {
		for {
			promptForPi()
		}
	}
	promptForPi()
}

func promptForPi() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your digits: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	correctDigits := PI[*startingDigit+1:*startingDigit+1+len(text)]
	if correctDigits == text {
		fmt.Print("Correct! - ")
		color.Green(text)
	} else {
		fmt.Print("Oops! It's actually: ")
		highlightCorrect(text)
		fmt.Print(", you entered: ")
		highlightIncorrect(text)
		fmt.Println()
	}
}

func highlightIncorrect(text string) {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	for i, char := range text {
		expected := PI[*startingDigit+1+i]
		if string(char) != string(expected) {
			fmt.Printf("%s", red(string(char)))
		} else {
			fmt.Printf("%s", yellow(string(char)))
		}
	}
}

func highlightCorrect(text string) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	for i, char := range text {
		expected := PI[*startingDigit+1+i]
		if string(char) != string(expected) {
			fmt.Printf("%s", green(string(expected)))
		} else {
			fmt.Printf("%s", yellow(string(expected)))
		}
	}
}
