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
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(number * 2)
	}
}
