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

func find_range_crabs_position(crabs_position []int) []int {
	min_position := 9999
	max_position := -1

	for i := 0; i < len(crabs_position); i++ {
		if crabs_position[i] < min_position {
			min_position = crabs_position[i]
		}
		if crabs_position[i] > max_position {
			max_position = crabs_position[i]
		}
	}
	return []int{min_position, max_position}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	crabs_position := array_str_to_int(strings.Split(scanner.Text(), ","))
	range_crabs_position := find_range_crabs_position(crabs_position)
	total_fuel := 9999999
	current_total_fuel := 0
	pos := 0

	for i := range_crabs_position[0]; i < range_crabs_position[1]; i++ {
		for j := 0; j < len(crabs_position); j++ {
			current_total_fuel += abs(i - crabs_position[j])
			if current_total_fuel > total_fuel {
				break
			}
		}
		if current_total_fuel < total_fuel {
			total_fuel = current_total_fuel
			pos = i
		}
		current_total_fuel = 0
	}

	fmt.Println("Total fuel:")
	fmt.Println(total_fuel)
	fmt.Println("Pos:")
	fmt.Println(pos)
}
