package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input"

func main() {
	answer1 := compute(12, 2)
	log.Printf("Answer #1 :: %d", answer1) // 3562624

	isFound, noun, verb := computeNounVerbForOutput(19690720)

	if isFound == true {
		log.Printf("Answer #2 :: %d", 100*noun+verb) // 8298
	}
}

func computeNounVerbForOutput(expectedOutput int) (bool, int, int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			answer := compute(noun, verb)

			if answer == expectedOutput {
				return true, noun, verb
			}
		}
	}

	return false, 0, 0
}

func compute(noun int, verb int) int {
	memory, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	memory[1] = noun
	memory[2] = verb

	// for answer 1
	for i := 0; i < len(memory)/4; i++ {
		opcodePosition := i * 4

		opcode := memory[opcodePosition]
		inputIndex1 := memory[opcodePosition+1]
		inputIndex2 := memory[opcodePosition+2]
		outputIndex := memory[opcodePosition+3]

		switch opcode {
		case 1:
			memory[outputIndex] = memory[inputIndex1] + memory[inputIndex2]
		case 2:
			memory[outputIndex] = memory[inputIndex1] * memory[inputIndex2]
		case 99:
			break
		}
	}

	return memory[0]
}

func getInput() ([]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	input, err := convertStringToArrayOfNumbers(line)

	if err != nil {
		return nil, err
	}

	return input, scanner.Err()
}

func convertStringToArrayOfNumbers(str string) ([]int, error) {
	var result []int

	for _, n := range strings.Split(str, ",") {
		num, err := strconv.Atoi(n)

		if err != nil {
			return nil, err
		}

		result = append(result, num)
	}

	return result, nil
}
