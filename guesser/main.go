package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
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

	guesses := 0
	guess := 10
	won := false
	//s := bufio.NewScanner(stdout)
	//s.Scan()
	//fmt.Println(s.Text())
	for won == false {
		guesses++
		go func() {
			//fmt.Println(guess)
			io.WriteString(stdin, fmt.Sprintf("%d\n", guess))
		}()
		mod := RecognizeInput(stdout)
		if mod == 0 {
			won = true
		}
		guess += mod
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal("wait failure: ", err)
	}

	fmt.Printf("%d guesses\n", guesses)
}

func RecognizeInput(reader io.Reader) int {
	m := []struct {
		A string
		Z int
	}{
		{"Correct", 0},
		{"Too high", -1},
		{"Too low", 1},
	}

	r := struct {
		A string
		Z int
	}{}

	b := bufio.NewReader(reader)
	for {
		c, _, err := b.ReadRune()
		if err != nil {
			if err == io.EOF {
				return 0
			}
			log.Fatal(err)
		}
		//fmt.Print(string(c))
		r.A = r.A + string(c)
		for _, i := range m {
			if strings.Contains(r.A, i.A) {
				return i.Z
			}
		}
	}
}
