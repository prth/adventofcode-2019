package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const inputFilePath = "input"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	orbitTree := convertPathsInputToOrbitTree(input)

	answer1 := computeTotalOrbitsCount(orbitTree, input)
	answer2 := computeShortestPathBetweenYouAndSanta(orbitTree)

	log.Printf("Answer #1 :: %d", answer1) // 162816
	log.Printf("Answer #2 :: %d", answer2) // 304
}

func convertPathsInputToOrbitTree(paths []string) OrbitTree {
	orbitTree := newOrbitTree("COM")

	for _, path := range paths {
		parentObjectKey, childObjectKey := getObjectKeysFromPathString(path)

		orbitTree.addBranch(parentObjectKey, childObjectKey)
	}

	return orbitTree
}

func computeTotalOrbitsCount(orbitTree OrbitTree, paths []string) int {
	return orbitTree.computeIndirectOrbitsCount() + len(paths)
}

func computeShortestPathBetweenYouAndSanta(orbitTree OrbitTree) int {
	yourPathToCOM := orbitTree.getPathFromNodeToRoot("YOU")
	santaPathToCOM := orbitTree.getPathFromNodeToRoot("SAN")

	// get the first intersection of two paths,
	// and use the indices in each path to compute path between you and santa
	for yourPathNodeIndex, yourPathNodeKey := range yourPathToCOM {
		for santaPathNodeIndex, santaPathNodeKey := range santaPathToCOM {
			if yourPathNodeKey == santaPathNodeKey {
				return yourPathNodeIndex + santaPathNodeIndex - 2
			}
		}
	}

	return -1
}

func getObjectKeysFromPathString(orbitString string) (string, string) {
	keys := strings.Split(orbitString, ")")

	if len(keys) != 2 {
		log.Panicf("Invalid orbit string %s", orbitString)
	}

	return keys[0], keys[1]
}

func getInput() ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input, scanner.Err()
}
