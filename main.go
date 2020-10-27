package main

import (
	"bufio"
	"flag"
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

	min := flag.Int("min", 1, "minimal secret")
	max := flag.Int("max", 10, "maximal secret")
	flag.Parse()
	if *max <= *min {
		log.Fatal("maximal secret should exceed minimal secret")
	}

	secret := rand.Intn(*max-*min+1) + *min
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number: ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Print("unable to read the line: ", err)
			continue
		}

		line = strings.TrimSpace(line)
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Print("unable to parse the number: ", err)
			continue
		}
		if number < *min {
			log.Print("number is too small")
			continue
		}
		if number > *max {
			log.Print("number is too large")
			continue
		}
		if number != secret {
			log.Print("you guessed wrong")
			continue
		}

		fmt.Println("You guessed right!")
		break
	}
}
