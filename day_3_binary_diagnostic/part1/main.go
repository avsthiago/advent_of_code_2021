package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculate_gamma(number_of_ones_per_position []int, total_lines int) []string {
	gamma := make([]string, 12)
	for i, value := range number_of_ones_per_position {
		if (total_lines - value) > (total_lines / 2) {
			gamma[i] = "1"
		} else {
			gamma[i] = "0"
		}
	}
	return gamma

}

func calculate_epsilon(number_of_ones_per_position []int, total_lines int) []string {
	epsilon := make([]string, 12)
	for i, value := range number_of_ones_per_position {
		if (total_lines - value) < (total_lines / 2) {
			epsilon[i] = "1"
		} else {
			epsilon[i] = "0"
		}
	}
	return epsilon
}

func rate_list_to_int(rate []string) int {
	rate_string := strings.Join(rate, "")
	rate_int, _ := strconv.ParseInt(rate_string, 2, 64)
	return int(rate_int)
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	number_of_ones_per_position := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	total_lines := 0

	for scanner.Scan() {
		current_read := strings.Split(scanner.Text(), "")

		for i, value := range current_read {
			if value == "1" {
				number_of_ones_per_position[i] += 1
			}
		}
		total_lines += 1
	}

	gamma := rate_list_to_int(calculate_gamma(number_of_ones_per_position[:], total_lines))
	epsilon := rate_list_to_int(calculate_epsilon(number_of_ones_per_position[:], total_lines))

	power_consumption := gamma * epsilon

	fmt.Println("Power consumption:")
	fmt.Println(power_consumption)
}
