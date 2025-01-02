package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me a name")
		return
	}
	name := args[0]

	emotions := []string{"happy", "terrible", "Sad", "Jolly", "verry happy"}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed)) // Create a new random generator with a known seed
	fmt.Printf("%s feeling today: %s\n", name, emotions[r.Intn(len(emotions))])

	
}
