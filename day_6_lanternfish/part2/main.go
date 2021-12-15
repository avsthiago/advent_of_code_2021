package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func array_str_to_int(array []string) []int {
	array_int := make([]int, 0)
	for _, value := range array {
		value_int, _ := strconv.Atoi(value)
		array_int = append(array_int, value_int)
	}
	return array_int
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	days := 256
	new_fishes := 0

	scanner.Scan()
	fishes := array_str_to_int(strings.Split(scanner.Text(), ","))
	fishes_multiplier := make([][]int, 0)

	for _, fish := range fishes {
		fishes_multiplier = append(fishes_multiplier, []int{fish, 0})
	}

	for i := 0; i < days; i++ {
		for j := 0; j < len(fishes_multiplier); j++ {
			if fishes_multiplier[j][0] == 0 {
				new_fishes += 1 + fishes_multiplier[j][1]
				fishes_multiplier[j][0] = 6
			} else {
				fishes_multiplier[j][0] -= 1
			}
		}

		if new_fishes > 0 {
			fishes_multiplier = append(fishes_multiplier, []int{8, new_fishes - 1})
		}
		new_fishes = 0
	}

	fmt.Println("Total fishes:")
	tot_fishes := 0
	for i := 0; i < len(fishes_multiplier); i++ {
		tot_fishes += fishes_multiplier[i][1] + 1
	}

	fmt.Println(tot_fishes)
}
