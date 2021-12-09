package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func find_oxygen_generator_rating(reads [][]string) []string {
	position := 0
	oxygen_generator_rating_reads := reads

	for there_are_more_reads(oxygen_generator_rating_reads) {
		most_frequent_bit := find_most_frequent_bit(position, oxygen_generator_rating_reads)
		oxygen_generator_rating_reads = filter_reads(oxygen_generator_rating_reads, most_frequent_bit, position)
		position += 1
	}

	return oxygen_generator_rating_reads[0]
}

func find_co2_scrubber_rating(reads [][]string) []string {
	position := 0
	co2_scrubber_rating_reads := reads

	for there_are_more_reads(co2_scrubber_rating_reads) {
		most_frequent_bit := find_most_frequent_bit(position, co2_scrubber_rating_reads)
		less_frequent_bit := less_frequent_bit(most_frequent_bit)
		co2_scrubber_rating_reads = filter_reads(co2_scrubber_rating_reads, less_frequent_bit, position)
		position += 1
	}

	return co2_scrubber_rating_reads[0]
}

func less_frequent_bit(most_frequent_bit string) string {
	if most_frequent_bit == "0" {
		return "1"
	} else {
		return "0"
	}
}

func find_most_frequent_bit(position int, reads [][]string) string {
	count_ones := 0
	for _, read := range reads {
		if read[position] == "1" {
			count_ones += 1
		}
	}
	if float32(count_ones) >= (float32(len(reads)) / 2.0) {
		return "1"
	} else {
		return "0"
	}
}

func filter_reads(reads [][]string, bit_to_filter string, position int) [][]string {
	filtered_list := make([][]string, 0)

	for _, read := range reads {
		if read[position] == bit_to_filter {
			filtered_list = append(filtered_list, read)
		}
	}

	return filtered_list
}

func there_are_more_reads(reads [][]string) bool {
	return len(reads) > 1
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

	reads := make([][]string, 0)

	for scanner.Scan() {
		current_read := strings.Split(scanner.Text(), "")
		reads = append(reads, current_read)
	}

	oxygen_generator_rating := rate_list_to_int(find_oxygen_generator_rating(reads))
	co2_scrubber_rating := rate_list_to_int(find_co2_scrubber_rating(reads))

	fmt.Println("Oxygen generator rating:")
	fmt.Println(oxygen_generator_rating)
	fmt.Println("CO2 scrubber rating:")
	fmt.Println(co2_scrubber_rating)
	fmt.Println("Life support rating:")
	fmt.Println(co2_scrubber_rating * oxygen_generator_rating)
}
