package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	previous := 99999
	total_increased := 0

	for scanner.Scan() {
		current_read, _ := strconv.Atoi(scanner.Text())
		if current_read > previous {
			total_increased += 1
		}
		previous = current_read
	}
	fmt.Println(total_increased)
}
