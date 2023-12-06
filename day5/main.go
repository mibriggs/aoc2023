package main

import (
	"Misc/aoc2023/day5/structs"
	"Misc/aoc2023/day5/utils"
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func main() {
	answer1 := puzzle1("input.txt")
	fmt.Printf("The answer to puzzle 1 is: %d\n", answer1)

	answer2 := puzzle2("input.txt")
	fmt.Printf("The answer to puzzle 2 is: %d\n", answer2)
}

func puzzle1(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seedIds := []int{}
	mapKey := ""
	rangeMap := map[string][]structs.Range{}
	currentRanges := []structs.Range{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentRanges) != 0 && mapKey != "" {
				val, exists := rangeMap[mapKey]
				if exists {
					rangeMap[mapKey] = append(val, currentRanges...)
				} else {
					rangeMap[mapKey] = currentRanges
				}
				currentRanges = []structs.Range{}
			}
		} else {
			lineArr := strings.Split(line, ":")
			lineLength := len(lineArr)

			if lineLength > 1 && lineArr[0] == "seeds" { // seeds line
				seedIds = utils.GetSeedIds(strings.Fields(lineArr[1]))
			} else if lineLength > 1 && lineArr[0] != "seeds" { //mapping line
				mapKey = utils.GetMapKey(line)
			} else { // range line
				currentRanges = append(currentRanges, structs.CreateRange(line))
			}
		}

	}

	seeds := utils.ContructSeedsArray(seedIds, rangeMap)
	closestSeed, err := utils.GetSeedWithSmallestLocation(seeds)
	shared.PanicIfError(err)

	err = scanner.Err()
	shared.PanicIfError(err)

	return closestSeed.Loc
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seedIds := []structs.SeedRange{}
	mapKey := ""
	rangeMap := map[string][]structs.Range{}
	currentRanges := []structs.Range{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentRanges) != 0 && mapKey != "" {
				val, exists := rangeMap[mapKey]
				if exists {
					rangeMap[mapKey] = append(val, currentRanges...)
				} else {
					rangeMap[mapKey] = currentRanges
				}
				currentRanges = []structs.Range{}
			}
		} else {
			lineArr := strings.Split(line, ":")
			lineLength := len(lineArr)

			if lineLength > 1 && lineArr[0] == "seeds" { // seeds line
				seedIds = utils.GetSeedIdsV2(strings.Fields(lineArr[1]))
			} else if lineLength > 1 && lineArr[0] != "seeds" { //mapping line
				mapKey = utils.GetMapKey(line)
			} else { // range line
				currentRanges = append(currentRanges, structs.CreateRange(line))
			}
		}

	}

	closestSeed := structs.Seed{Loc: math.MaxInt}

	for _, seedRange := range seedIds {
		for i := 0; i < seedRange.Range; i++ {
			newSeedId := seedRange.Start + i
			newSeed := structs.CreateSeed(newSeedId)

			oldSourceId := newSeedId
			for _, stop := range utils.TraversalOrder {
				ranges := rangeMap[stop]
				id := utils.DoesSourceExistInRange(oldSourceId, ranges)
				newId := id
				if id == -1 {
					newId = oldSourceId
				}
				newSeed.SetParam(stop, newId)
				oldSourceId = newId
			}

			if newSeed.IsCloser(closestSeed) {
				closestSeed = newSeed
			}
		}

	}
	return closestSeed.Loc
}
