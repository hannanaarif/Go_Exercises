package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxturns = 5
	usage    = `Welcome to the Lucky Number Game! 🎲
This game will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	winMessage := []string{
		"🎉 You win! Awesome job!",
		"🥳 Victory is yours!",
		"🏆 Congratulations, you nailed it!",
	}

	lossMessage := []string{
		"❌ You lost! Try again!",
		"😢 Not this time. Better luck next time!",
		"💔 Oh no, you missed it!",
	}

	fmt.Println(fmt.Sprintf(usage, maxturns))
	// fmt.Println(os.Args)
	// fmt.Println("os.Args[0]",os.Args[0])
	// fmt.Println("os.Args[1]",os.Args[1:])

	args := os.Args[1:]

	if len(args) <= 1 {
		fmt.Println("Please provide a number")
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess1, err := strconv.Atoi(args[1])

	if guess < 0  || guess1 < 0 {
		fmt.Println("Please provide a positive number")
		return
	}

	if err != nil {
		fmt.Println("Not a number")
		return
	}

	for turn := 0; turn < maxturns; turn++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("Turn %d: %d\n", turn, n)
		if n == guess || n == guess1 {
			if turn == 0 {
				fmt.Println("🎉 You win! very first time congratulation 🎉")
				return
			}else{
			fmt.Println("✅", winMessage[rand.Intn(len(winMessage))])
			return
			}

		}
	}
	fmt.Println("❌", lossMessage[rand.Intn(len(lossMessage))])
}
