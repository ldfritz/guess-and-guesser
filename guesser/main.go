package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	binPath := os.Args[1]
	cmd := exec.Command(binPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		log.Fatal("run failure: ", err)
	}

	guess := 10
	scanner := bufio.NewScanner(stdout)
	won := false
	var winMsg = regexp.MustCompile(`^Correct`)
	for won == false {
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(input)
		switch {
		case input == "Too high":
			guess--
		case input == "Too low":
			guess++
		case winMsg.MatchString(input):
			won = true
		default:
			go func() {
				fmt.Println(guess)
				io.WriteString(stdin, fmt.Sprintf("%d\n", guess))
			}()
		}
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal("wait failure: ", err)
	}
}
