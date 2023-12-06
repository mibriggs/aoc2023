package utils

import (
	"Misc/aoc2023/day5/structs"
	"errors"
	"sort"
	"strconv"
	"strings"
)

var TraversalOrder = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func GetSeedIds(seeds []string) []int {
	ids := []int{}
	for _, currentId := range seeds {
		intId, err := strconv.Atoi(currentId)
		if err == nil {
			ids = append(ids, intId)
		}
	}
	return ids
}

func GetSeedIdsV2(seeds []string) []structs.SeedRange {
	ids := []structs.SeedRange{}
	currentSeedRange := structs.SeedRange{}
	for indx, currentId := range seeds {
		intId, _ := strconv.Atoi(currentId)

		if indx%2 != 0 {
			currentSeedRange.Range = intId
			ids = append(ids, currentSeedRange)
			currentSeedRange = structs.SeedRange{}
		}
		currentSeedRange.Start = intId
	}
	return ids
}

func GetSeedWithSmallestLocation(seeds []structs.Seed) (*structs.Seed, error) {
	if len(seeds) == 0 {
		return nil, errors.New("list is empty")
	}
	sort.Sort(structs.SortByLocation(seeds))
	return &seeds[0], nil
}

func SetMapKey(nameLine string) string {
	arr := strings.Fields(nameLine)
	names := strings.Split(arr[0], "-to-")
	name := names[1]
	name = strings.Trim(name, " ")
	return strings.ToLower(name)
}

func DoesSourceExistInRange(sourceId int, ranges []structs.Range) int {
	for _, currRange := range ranges {
		destId := currRange.GetDestId(sourceId)
		if destId != -1 {
			return destId
		}
	}
	return -1
}

func ContructSeedsArray(seedIds []int, rangeMap map[string][]structs.Range) []structs.Seed {
	seeds := []structs.Seed{}
	for _, seedId := range seedIds {
		seed := structs.CreateSeed(seedId)
		oldSourceId := seedId
		for _, stop := range TraversalOrder {
			ranges := rangeMap[stop]
			id := DoesSourceExistInRange(oldSourceId, ranges)
			newId := id
			if id == -1 {
				newId = oldSourceId
			}
			seed.SetParam(stop, newId)
			oldSourceId = newId
		}
		seeds = append(seeds, seed)
	}
	return seeds
}
