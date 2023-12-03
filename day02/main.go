package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_RED_CUBES_IN_GAME = 12
const MAX_GREEN_CUBES_IN_GAME = 13
const MAX_BLUE_CUBES_IN_GAME = 14

type Round struct {
	blueCubes  int
	redCubes   int
	greenCubes int
}

type Game struct {
	id     int
	rounds []Round
}

func extractNumberOfCubes(cubeCountString string) int {
	numberRegex := regexp.MustCompile(`\d+`)

	countAsString := numberRegex.FindString(cubeCountString)
	count, _ := strconv.Atoi(countAsString)

	return count
}

func parseGame(gameAsString string) Game {
	numberRegex := regexp.MustCompile(`\d+`)
	idAsString := numberRegex.FindString(gameAsString)
	id, _ := strconv.Atoi(idAsString)

	gameWithOnlyRounds := strings.Split(gameAsString, ":")[1]
	roundsAsStringArray := strings.Split(strings.Trim(gameWithOnlyRounds, " "), ";")

	var rounds []Round = []Round{}

	blueRegex := regexp.MustCompile(`\d+ blue`)
	greenRegex := regexp.MustCompile(`\d+ green`)
	redRegex := regexp.MustCompile(`\d+ red`)
	for _, round := range roundsAsStringArray {
		cubeCounts := strings.Split(strings.Trim(round, " "), ",")
		currentRound := Round{
			blueCubes:  0,
			greenCubes: 0,
			redCubes:   0,
		}
		for _, cubeCount := range cubeCounts {
			if blueRegex.MatchString(cubeCount) {
				currentRound.blueCubes = extractNumberOfCubes(cubeCount)
			}

			if greenRegex.MatchString(cubeCount) {
				currentRound.greenCubes = extractNumberOfCubes(cubeCount)
			}

			if redRegex.MatchString(cubeCount) {
				currentRound.redCubes = extractNumberOfCubes(cubeCount)
			}
		}

		rounds = append(rounds, currentRound)
	}

	return Game{
		id:     id,
		rounds: rounds,
	}
}

func solutionForPart1(games []Game) int {
	total := 0

	for _, game := range games {
		isGamePossible := true

		for _, round := range game.rounds {
			if round.blueCubes > MAX_BLUE_CUBES_IN_GAME || round.greenCubes > MAX_GREEN_CUBES_IN_GAME || round.redCubes > MAX_RED_CUBES_IN_GAME {
				isGamePossible = false
			}
		}

		if isGamePossible {
			total += game.id
		}
	}

	return total
}

func solutionForPart2(games []Game) int {
	sum := 0

	for _, game := range games {
		// We initialize to one so that 0 does not get ruin the multiplication
		redMax := 1
		blueMax := 1
		greenMax := 1

		for _, round := range game.rounds {
			if redMax < round.redCubes {
				redMax = round.redCubes
			}

			if blueMax < round.blueCubes {
				blueMax = round.blueCubes
			}

			if greenMax < round.greenCubes {
				greenMax = round.greenCubes
			}
		}

		sum += (redMax * blueMax * greenMax)
	}

	return sum
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Could not open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Build out the games array and parse the file to make this easier
	games := []Game{}
	for scanner.Scan() {
		gameAsString := scanner.Text()
		parsedGame := parseGame(gameAsString)
		games = append(games, parsedGame)
		fmt.Println("Parsed Game", parsedGame)
	}

	part1Solution := solutionForPart1(games)
	part2Solution := solutionForPart2(games)

	fmt.Println("Solution for Part 1", part1Solution)
	fmt.Println("Solution for Part 2", part2Solution)
}
