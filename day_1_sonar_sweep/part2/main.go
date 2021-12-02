package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mov(new int, p1 int, p2 int, p3 int) (int, int, int) {
	p3 = p2
	p2 = p1
	p1 = new

	return p1, p2, p3
}

func getScannedValue(scanner *bufio.Scanner) int {
	scanner.Scan()
	scannedValue, _ := strconv.Atoi(scanner.Text())
	return scannedValue
}

func getSum(p1 int, p2 int, p3 int) (int) {
	return p1 + p2 + p3
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	p1 := getScannedValue(scanner)
	p2 := getScannedValue(scanner)
	p3 := getScannedValue(scanner)

	previous := getSum(p1, p2, p3)
	total_increased := 0

	for scanner.Scan() {
		current_measurement, _ := strconv.Atoi(scanner.Text())

		p1, p2, p3 = mov(current_measurement, p1, p2, p3)

		current_sum_measurement := getSum(p1, p2, p3)

		if current_sum_measurement > previous {
			total_increased += 1
		}
		previous = current_sum_measurement
	}
	fmt.Println(total_increased)
}
