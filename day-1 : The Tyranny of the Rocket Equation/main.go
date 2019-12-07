package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	sumAnswer1 := 0
	sumAnswer2 := 0

	for _, mass := range input {
		fuel := getFuelRequired((mass))

		sumAnswer1 += fuel
		sumAnswer2 += fuel

		// now to calculate the fuel required for fuel...
		for fuel != 0 {
			fuel = getFuelRequired(fuel)
			sumAnswer2 += fuel
		}
	}

	log.Printf("Answer #1 :: %d", sumAnswer1) // 3330521
	log.Printf("Answer #2 :: %d", sumAnswer2) // 4992931
}

func getFuelRequired(mass int) int {
	fuel := (mass / 3) - 2

	if fuel < 0 {
		return 0
	}

	return fuel
}

func getInput() ([]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []int

	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line)
		input = append(input, num)
	}

	return input, scanner.Err()
}
