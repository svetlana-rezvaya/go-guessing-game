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
	reader := bufio.NewReader(os.Stdin)
	for {
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

		fmt.Println(number * 2)
	}
}
