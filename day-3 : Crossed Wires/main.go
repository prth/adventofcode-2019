package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input"

func main() {
	wire1Path, wire2Path, err := getInput()

	if err != nil {
		log.Fatal(err)
		return
	}

	wire1PathLines := convertWirePathToLines(wire1Path)
	wire2PathLines := convertWirePathToLines(wire2Path)

	closestIntManhattanPoint, closestIntPathPoint := getClosestIntersectionsBetweenWirePaths(wire1PathLines, wire2PathLines)

	if closestIntManhattanPoint != nil {
		log.Printf("Answer #1 :: %d", getManhattanDistance(closestIntManhattanPoint)) // 8015
	}

	if closestIntPathPoint != nil {
		answer2 := getPathDistance(closestIntPathPoint, wire1PathLines) + getPathDistance(closestIntPathPoint, wire2PathLines)
		log.Printf("Answer #1 :: %d", answer2) // 163676
	}
}

func getClosestIntersectionsBetweenWirePaths(wire1PathLines []line, wire2PathLines []line) (*point, *point) {
	var closestIntManhattanPoint *point
	var closestIntManhattanDist int

	var closestIntPathPoint *point
	var closestIntPathDist int
	var preVisitedPointDistances map[*point]int

	for _, wire1EachLine := range wire1PathLines {
		for _, wire2EachLine := range wire2PathLines {
			newIntPoint := getIntersectionBetweenTwoLines(wire1EachLine, wire2EachLine)

			if newIntPoint == nil || (newIntPoint.x == 0 && newIntPoint.y == 0) {
				continue
			}

			newManhattanDist := getManhattanDistance(newIntPoint)

			if closestIntManhattanPoint == nil || newManhattanDist < closestIntManhattanDist {
				closestIntManhattanPoint = newIntPoint
				closestIntManhattanDist = newManhattanDist
			}

			var newPathDist int

			if val, ok := preVisitedPointDistances[newIntPoint]; ok {
				newPathDist = val
			} else {
				newPathDist = getPathDistance(newIntPoint, wire1PathLines) + getPathDistance(newIntPoint, wire2PathLines)
			}

			if closestIntPathPoint == nil || newPathDist < closestIntPathDist {
				closestIntPathPoint = newIntPoint
				closestIntPathDist = newPathDist
			}
		}
	}

	return closestIntManhattanPoint, closestIntPathPoint
}

func getManhattanDistance(p *point) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func getPathDistance(targetPoint *point, wire1PathLines []line) int {
	distance := 0

	for _, eachLine := range wire1PathLines {
		if !isPointOnALine(targetPoint, eachLine) {
			distance += calculateLineDistance(eachLine)
		} else {
			return distance + calculateLineDistance(line{
				point1: eachLine.point1,
				point2: point{
					x: targetPoint.x,
					y: targetPoint.y,
				},
			})
		}
	}

	return distance
}

func calculateLineDistance(l line) int {
	return int(math.Sqrt(math.Pow(float64(l.point1.x-l.point2.x), 2) + math.Pow(float64(l.point1.y-l.point2.y), 2)))
}

func isPointOnALine(p *point, l line) bool {
	if isLineHorizontal(l) {
		return l.point1.y == p.y && inIntBetween(p.x, l.point1.x, l.point2.x)
	}

	return l.point1.x == p.x && inIntBetween(p.y, l.point1.y, l.point2.y)
}

func getIntersectionBetweenTwoLines(line1 line, line2 line) *point {
	// if line 1 is horizontal, check if line 2 crosses it by
	// checking if line2 x-cordinate falls between line 1
	if isLineHorizontal(line1) {
		if inIntBetween(line1.point1.y, line2.point1.y, line2.point2.y) &&
			inIntBetween(line2.point1.x, line1.point1.x, line1.point2.x) {
			return &point{
				x: line2.point1.x,
				y: line1.point1.y,
			}
		}
	} else {
		// if line 1 is vertical, check if line 2 crosses it by
		// checking if line2 y-cordinate falls between line 1
		if inIntBetween(line1.point1.x, line2.point1.x, line2.point2.x) &&
			inIntBetween(line2.point1.y, line1.point1.y, line1.point2.y) {
			return &point{
				x: line1.point1.x,
				y: line2.point1.y,
			}
		}
	}

	return nil
}

func inIntBetween(target int, num1 int, num2 int) bool {
	if num1 <= num2 {
		return (num1 <= target) && (target <= num2)
	}

	return (num2 <= target) && (target <= num1)
}

func isLineHorizontal(l line) bool {
	return l.point1.y == l.point2.y
}

func convertWirePathToLines(wirePath []string) []line {
	var pathLines []line

	currentPosition := point{
		x: 0,
		y: 0,
	}

	for _, linePath := range wirePath {
		direction, distance := getDirectionAndDistanceForALinePath(linePath)

		var newPosition point

		switch direction {
		case "R":
			newPosition = point{
				x: currentPosition.x + distance,
				y: currentPosition.y,
			}

		case "L":
			newPosition = point{
				x: currentPosition.x - distance,
				y: currentPosition.y,
			}

		case "U":
			newPosition = point{
				x: currentPosition.x,
				y: currentPosition.y + distance,
			}

		case "D":
			newPosition = point{
				x: currentPosition.x,
				y: currentPosition.y - distance,
			}
		}

		pathLines = append(pathLines, line{
			point1: currentPosition,
			point2: newPosition,
		})
		currentPosition = newPosition
	}

	return pathLines
}

func getDirectionAndDistanceForALinePath(path string) (string, int) {
	direction := string(path[0])
	distance, _ := strconv.Atoi(path[1:])

	return direction, distance
}

func getInput() ([]string, []string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	wire1Path := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	wire2Path := strings.Split(scanner.Text(), ",")

	return wire1Path, wire2Path, nil
}

type point struct {
	x int
	y int
}

type line struct {
	point1 point
	point2 point
}
