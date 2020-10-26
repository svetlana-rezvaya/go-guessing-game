package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(10) + 1
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number: ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}

		line = strings.TrimSpace(line)
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Print(err)
			continue
		}
		if number < 1 {
			log.Print("number is too small")
			continue
		}
		if number > 10 {
			log.Print("number is too large")
			continue
		}
		if number != secret {
			log.Print("you guessed wrong")
			continue
		}

		fmt.Println("you guessed right")
		break
	}
}
