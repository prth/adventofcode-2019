package main

import (
	"log"
)

const rangeMin = 264360
const rangeMax = 746325

type passwordCriteria int

const (
	pwdCriteriaOne passwordCriteria = 1
	pwdCriteriaTwo passwordCriteria = 2
)

func main() {
	validPasswordsAns1 := 0
	validPasswordsAns2 := 0

	for i := rangeMin; i <= rangeMax; i++ {
		if isPasswordValid(i, pwdCriteriaOne) {
			validPasswordsAns1++
		}

		if isPasswordValid(i, pwdCriteriaTwo) {
			validPasswordsAns2++
		}
	}

	log.Printf("Answer #1 :: %d", validPasswordsAns1)
	log.Printf("Answer #2 :: %d", validPasswordsAns2)
}

func isPasswordValid(password int, criteria passwordCriteria) bool {
	// It is a six-digit number.
	// The value is within the range given in your puzzle input.
	if password < rangeMin || password > rangeMax {
		return false
	}

	areAnyTwoAdjacentDigitsSame := false

	digits := getDigitsOfNumber(password)

	for i := 0; i < 5; i++ {
		// Going from left to right, the digits never decrease; return early if they do
		// we are checking from right to left
		if digits[i] < digits[i+1] {
			return false
		}

		if !areAnyTwoAdjacentDigitsSame && digits[i] == digits[i+1] {
			areAnyTwoAdjacentDigitsSame = true

			// for answer#2 password criteria two,
			// the two adjacent matching digits are not part of
			// a larger group of matching digits.
			// so, lets negate the flag if neighboring digits (right or left) are same
			if criteria == pwdCriteriaTwo {
				if i < 4 && digits[i+1] == digits[i+2] {
					areAnyTwoAdjacentDigitsSame = false
				}

				if i >= 1 && digits[i] == digits[i-1] {
					areAnyTwoAdjacentDigitsSame = false
				}
			}
		}
	}

	// return the state of the adjacent digits flag
	// we have already confirmed that the digits never decrease,
	// and if they did `false` got returned early
	return areAnyTwoAdjacentDigitsSame
}

func getDigitsOfNumber(number int) []int {
	var result []int

	for number > 10 {
		result = append(result, number%10)
		number = number / 10
	}

	result = append(result, number)

	return result
}
