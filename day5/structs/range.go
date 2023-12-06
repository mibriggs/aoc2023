package structs

import (
	"Misc/aoc2023/shared"
	"strconv"
	"strings"
)

type Range struct {
	Dest        int
	Source      int
	RangeLength int
}

func (currRange Range) IsEmpty() bool {
	return currRange.Dest == -1 && currRange.Source == -1 && currRange.RangeLength == -1
}

func (currRange Range) GetDestId(sourceId int) int {
	maxSource := currRange.Source + (currRange.RangeLength - 1)
	minSource := currRange.Source
	diff := sourceId - minSource
	if sourceId >= minSource && sourceId <= maxSource {
		return currRange.Dest + diff
	}
	return -1
}

func CreateRange(rangeString string) Range {
	items := strings.Fields(rangeString)
	newRange := Range{}
	for indx, item := range items {
		asInt, err := strconv.Atoi(item)
		shared.PanicIfError(err)
		if indx == 0 {
			newRange.Dest = asInt
		} else if indx == 1 {
			newRange.Source = asInt
		} else {
			newRange.RangeLength = asInt
		}
	}
	return newRange
}

func CreateEmptyRange() Range {
	return Range{
		Source:      -1,
		Dest:        -1,
		RangeLength: -1,
	}
}
