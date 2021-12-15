package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process_read(read string) []int {
	read_with_comma := strings.ReplaceAll(read, " -> ", ",")
	read_array_str := strings.Split(read_with_comma, ",")
	return array_str_to_int(read_array_str)
}

func array_str_to_int(array []string) []int {
	array_int := make([]int, 0)
	for _, value := range array {
		value_int, _ := strconv.Atoi(value)
		array_int = append(array_int, value_int)
	}
	return array_int
}

func find_map_size(reads [][]int) int {
	size := 0
	for _, reads_line := range reads {
		for _, read := range reads_line {
			if read > size {
				size = read
			}
		}
	}
	return size + 1
}

func init_map(size int) [][]int {
	_map := make([][]int, 0)
	line := make([]int, 0)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			line = append(line, 0)
		}
		_map = append(_map, line)
		line = nil
	}
	return _map
}

func find_vents(_map [][]int, reads [][]int) [][]int {
	for _, read := range reads {
		read = sort_points(read)
		range_x := get_range(read[0], read[2])
		range_y := get_range(read[1], read[3])
		if is_horizontal_or_vertical(read) {
			for i := range_x[0]; i <= range_x[1]; i++ {
				for j := range_y[0]; j <= range_y[1]; j++ {
					_map[j][i] += 1
				}
			}
		} else if is_diagonal(read) {
			if is_primary(read) {
				for i := range_x[0]; i <= range_x[1]; i++ {
					for j := range_y[0]; j <= range_y[1]; j++ {
						if is_diagonal([]int{range_x[0], range_y[0], i, j}) {
							_map[j][i] += 1
						}
					}
				}
			} else {
				for i := read[0]; i <= read[2]; i++ {
					for j := read[1]; j >= read[3]; j-- {
						if is_diagonal([]int{read[0], read[1], i, j}) && !is_primary([]int{read[0], read[1], i, j}) {
							_map[j][i] += 1
						}
					}
				}
			}
		}
	}
	return _map
}

func is_horizontal_or_vertical(read []int) bool {
	return read[0] == read[2] || read[1] == read[3]
}

func is_diagonal(read []int) bool {
	return abs(read[0]-read[2]) == abs(read[1]-read[3])
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func sort_points(read []int) []int {
	if read[0] > read[2] {
		return []int{read[2], read[3], read[0], read[1]}
	}
	return read
}

func is_primary(read []int) bool {
	return read[1] < read[3]
}

func get_range(v1 int, v2 int) []int {
	if v1 > v2 {
		return []int{v2, v1}
	}
	return []int{v1, v2}
}

func find_num_danger_zones(_map [][]int) int {
	num_danger_zones := 0
	for _, line := range _map {
		for _, val_zone := range line {
			if val_zone >= 2 {
				num_danger_zones += 1
			}
		}
	}
	return num_danger_zones
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	reads := make([][]int, 0)

	for scanner.Scan() {
		current_read := scanner.Text()
		reads = append(reads, process_read(current_read))

	}

	size := find_map_size(reads)

	_map := init_map(size)
	map_with_vents := find_vents(_map, reads)
	num_danger_zones := find_num_danger_zones(map_with_vents)

	fmt.Println("num_danger_zones")
	fmt.Println(num_danger_zones)
}
