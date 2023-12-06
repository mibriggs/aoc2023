package structs

import "strings"

type SourceDestMapping struct {
	Name   string
	Ranges []Range
}

func CreateEmptyMapping() SourceDestMapping {
	return SourceDestMapping{
		Name:   "",
		Ranges: nil,
	}
}

func (mapping SourceDestMapping) IsEmpty() bool {
	return mapping.Name == "" && mapping.Ranges == nil
}

func (mapping *SourceDestMapping) SetName(nameLine string) {
	arr := strings.Fields(nameLine)
	mapping.Name = arr[0]
}

func (mapping *SourceDestMapping) Push(newRange Range) {
	if mapping.Ranges == nil {
		mapping.Ranges = []Range{newRange}
	} else {
		mapping.Ranges = append(mapping.Ranges, newRange)
	}
}
