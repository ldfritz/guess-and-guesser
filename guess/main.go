package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := 5
	guess := -1
	tries := 0
	reader := bufio.NewReader(os.Stdin)
	for guess != answer {
		tries++
		fmt.Print("Guess a number:\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}
		if guess > answer {
			fmt.Println("Too high")
		}
		if guess < answer {
			fmt.Println("Too low")
		}
	}
	fmt.Printf("Correct! (%d tries)\n", tries)
}
