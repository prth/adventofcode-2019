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
	answer1 := compute(1)
	log.Printf("Answer #1 :: %d", answer1) // 7286649

	answer2 := compute(5)
	log.Printf("Answer #2 :: %d", answer2) // 15724522
}

func compute(input int) int {
	memory, err := getInput()
	var output []int

	if err != nil {
		log.Fatal(err)
	}

	// for answer 1
	for instructionPosition := 0; instructionPosition < len(memory); {
		instruction := memory[instructionPosition]

		opcode, parameterModes := decodeInstruction(instruction)
		opcodeParamsLength := getOpcodeParamsLength(opcode)

		var paramIndices []int

		// collect params for the opcode
		for j := 0; j < opcodeParamsLength; j++ {
			var inputIndex int
			if parameterModes[j] == 0 {
				inputIndex = memory[instructionPosition+j+1]
			} else {
				inputIndex = instructionPosition + j + 1
			}

			paramIndices = append(paramIndices, inputIndex)
		}

		switch opcode {
		// addition
		case 1:
			memory[paramIndices[2]] = memory[paramIndices[0]] + memory[paramIndices[1]]

		// multiplication
		case 2:
			memory[paramIndices[2]] = memory[paramIndices[0]] * memory[paramIndices[1]]

		// replace with input
		case 3:
			memory[paramIndices[0]] = input

			// output
		case 4:
			output = append(output, memory[paramIndices[0]])

		case 7:
			if memory[paramIndices[0]] < memory[paramIndices[1]] {
				memory[paramIndices[2]] = 1
			} else {
				memory[paramIndices[2]] = 0
			}

		case 8:
			if memory[paramIndices[0]] == memory[paramIndices[1]] {
				memory[paramIndices[2]] = 1
			} else {
				memory[paramIndices[2]] = 0
			}

		case 99:
			// return the last output as diagnostic code
			return output[len(output)-1]
		}

		// go to the next instruction
		switch opcode {
		// jump-if-true
		case 5:
			if memory[paramIndices[0]] != 0 {
				instructionPosition = memory[paramIndices[1]]
				continue
			}

		// jump-if-false
		case 6:
			if memory[paramIndices[0]] == 0 {
				instructionPosition = memory[paramIndices[1]]
				continue
			}

		}

		// skip to next instruction
		instructionPosition = instructionPosition + opcodeParamsLength + 1
	}

	return output[len(output)-1]
}

func decodeInstruction(instruction int) (int, []int) {
	// get last 2 digits as the opcode
	opcode := instruction % 100

	var parameterModes []int

	// get the intruction int after exctracting the opcode
	paramModeInst := instruction / 100

	for paramModeInst >= 10 {
		parameterModes = append(parameterModes, paramModeInst%10)
		paramModeInst = paramModeInst / 10
	}

	parameterModes = append(parameterModes, paramModeInst)

	// prefill default param modes of 0, if not available in instruction
	if len(parameterModes) < 3 {
		for i := 0; i <= 3-len(parameterModes); i++ {
			parameterModes = append(parameterModes, 0)
		}
	}

	return opcode, parameterModes
}

func getOpcodeParamsLength(opcode int) int {
	switch opcode {
	case 1, 2, 7, 8:
		return 3
	case 5, 6:
		return 2
	case 3, 4:
		return 1
	}

	return 0
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
