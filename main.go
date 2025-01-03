package main

import (
	"fmt"
	"log"
	"math/rand"
)

func neighboursCount(i int, j int, universe [][]string) int {
	var aliveNeighborsCount int
	var universeSize = len(universe)

	newJ := eastNeighbour(j, universeSize)
	if isNeighbourAlive(i, newJ, universe) {
		aliveNeighborsCount++
	}

	newJ = westNeighbour(j, universeSize)
	if isNeighbourAlive(i, newJ, universe) {
		aliveNeighborsCount++
	}

	newI := northNeighbour(i, universeSize)
	if isNeighbourAlive(newI, j, universe) {
		aliveNeighborsCount++
	}

	newI = southNeighbour(i, universeSize)
	if isNeighbourAlive(newI, j, universe) {
		aliveNeighborsCount++
	}

	newI, newJ = northWestNeighbour(i, j, universeSize)
	if isNeighbourAlive(newI, newJ, universe) {
		aliveNeighborsCount++
	}

	newI, newJ = southWestNeighbour(i, j, universeSize)
	if isNeighbourAlive(newI, newJ, universe) {
		aliveNeighborsCount++
	}

	newI, newJ = northEastNeighbour(i, j, universeSize)
	if isNeighbourAlive(newI, newJ, universe) {
		aliveNeighborsCount++
	}

	newI, newJ = southEastNeighbour(i, j, universeSize)
	if isNeighbourAlive(newI, newJ, universe) {
		aliveNeighborsCount++
	}

	return aliveNeighborsCount
}

func isNeighbourAlive(i, j int, universe [][]string) bool {
	return universe[i][j] != " "
}

func northWestNeighbour(i, j, universeSize int) (int, int) {
	if i == 0 {
		i = universeSize - 1
	} else {
		i--
	}

	if j == 0 {
		j = universeSize - 1
	} else {
		j--
	}

	return i, j
}

func westNeighbour(j, universeSize int) int {
	if j == 0 {
		return universeSize - 1
	}

	return j - 1
}

func southWestNeighbour(i, j, universeSize int) (int, int) {
	if i == universeSize-1 {
		i = 0
	} else {
		i++
	}

	if j == 0 {
		j = universeSize - 1
	} else {
		j--
	}

	return i, j
}

func northNeighbour(i, universeSize int) int {
	if i == 0 {
		return universeSize - 1
	}

	return i - 1
}

func southNeighbour(i, universeSize int) int {
	if i == universeSize-1 {
		return 0
	}

	return i + 1
}

func northEastNeighbour(i, j, universeSize int) (int, int) {
	if i == universeSize-1 {
		i = 0
	} else {
		i++
	}

	if j == universeSize-1 {
		j = 0
	} else {
		j++
	}

	return i, j
}

func eastNeighbour(j, universeSize int) int {
	if j == universeSize-1 {
		return 0
	}

	return j + 1
}

func southEastNeighbour(i, j, universeSize int) (int, int) {
	if i == 0 {
		i = universeSize - 1
	} else {
		i--
	}

	if j == universeSize-1 {
		j = 0
	} else {
		j++
	}

	return i, j
}

func initUniverse(universeSize int, seed int64) [][]string {
	r := rand.New(rand.NewSource(seed))

	universe := make([][]string, universeSize)

	for i := 0; i < universeSize; i++ {
		universe[i] = make([]string, universeSize)
		for j := 0; j < universeSize; j++ {
			if r.Intn(2) == 1 {
				universe[i][j] = "O"
				continue
			}
			universe[i][j] = " "
		}
	}

	return universe
}

func printUniverse(universe [][]string) {
	var universeSize = len(universe)

	for i := 0; i < universeSize; i++ {
		for j := 0; j < universeSize; j++ {
			fmt.Printf("%s", universe[i][j])
		}
		fmt.Println()
	}
}

func emptyUniverse(universeSize int) [][]string {
	universe := make([][]string, universeSize)

	for i := 0; i < universeSize; i++ {
		universe[i] = make([]string, universeSize)
		for j := 0; j < universeSize; j++ {
			universe[i][j] = " "
		}
	}
	return universe
}

func buildGenerations(universe [][]string) [][]string {
	var universeSize = len(universe)
	nextGeneration := emptyUniverse(universeSize)

	//check current cell
	for i := 0; i < universeSize; i++ {
		for j := 0; j < universeSize; j++ {
			count := neighboursCount(i, j, universe)
			if universe[i][j] == "O" {
				if count == 2 || count == 3 {
					nextGeneration[i][j] = "O"
				} else {
					nextGeneration[i][j] = " "
				}
			}
			if universe[i][j] == " " {
				if count == 3 {
					nextGeneration[i][j] = "O"
				} else {
					nextGeneration[i][j] = " "
				}
			}
		}
	}
	return nextGeneration
}

func main() {
	var universeSize int
	var seed int64
	var generationsCount int
	_, err := fmt.Scan(&universeSize, &seed, &generationsCount)
	if err != nil {
		log.Fatal(err)
	}
	if universeSize <= 0 || seed < 0 {
		log.Fatal("universeSize and seed must be greater than zero")
	}

	universe := initUniverse(universeSize, seed)
	for i := 0; i < generationsCount; i++ {
		universe = buildGenerations(universe)
	}
	printUniverse(universe)
}
