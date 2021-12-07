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

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	aim := 0
	vertical_position := 0
	horizontal_position := 0

	for scanner.Scan() {
		current_read := strings.Split(scanner.Text(), " ")
		movement := current_read[0]
		ammount, _ := strconv.Atoi(current_read[1])

		switch movement {
		case "forward":
			vertical_position += ammount
			horizontal_position += aim * ammount
		case "up":
			aim -= ammount
		case "down":
			aim += ammount
		}
	}

	fmt.Println("vertical_position:")
	fmt.Println(vertical_position)
	fmt.Println("horizontal_position:")
	fmt.Println(horizontal_position)
	fmt.Println("aim:")
	fmt.Println(horizontal_position)
	fmt.Println("multiplication vertical_position x horizontal_position:")
	fmt.Println(horizontal_position * vertical_position)
}
