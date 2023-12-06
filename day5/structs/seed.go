package structs

type Seed struct {
	Id         int
	Soil       int
	Fertilizer int
	Water      int
	Light      int
	Temp       int
	Humidity   int
	Loc        int
}

type SortByLocation []Seed

func (a SortByLocation) Len() int           { return len(a) }
func (a SortByLocation) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLocation) Less(i, j int) bool { return a[i].Loc < a[j].Loc }

// Mutates the current Seed to update its values
func (seed *Seed) SetParam(stop string, id int) {
	if stop == "soil" {
		seed.Soil = id
	} else if stop == "fertilizer" {
		seed.Fertilizer = id
	} else if stop == "water" {
		seed.Water = id
	} else if stop == "light" {
		seed.Light = id
	} else if stop == "temperature" {
		seed.Temp = id
	} else if stop == "humidity" {
		seed.Humidity = id
	} else {
		seed.Loc = id
	}
}

// Does the current Seed have a smaller location that the other seed?
func (currSeed Seed) IsCloser(otherSeed Seed) bool {
	return currSeed.Loc <= otherSeed.Loc
}

// Is the current Seed empty?
func (seed Seed) IsEmpty() bool {
	return seed.Id == -1 &&
		seed.Loc == -1 &&
		seed.Soil == -1 &&
		seed.Fertilizer == -1 &&
		seed.Water == -1 &&
		seed.Light == -1 &&
		seed.Temp == -1 &&
		seed.Humidity == -1
}

// Creates new Seed with given Seed Id and everything else -1
func CreateSeed(seedId int) Seed {
	return Seed{
		Id:         seedId,
		Soil:       -1,
		Fertilizer: -1,
		Water:      -1,
		Light:      -1,
		Temp:       -1,
		Humidity:   -1,
		Loc:        -1,
	}
}

// Creates a Seed with all values set to -1
func CreateEmptySeed() Seed {
	return Seed{
		Id:         -1,
		Soil:       -1,
		Fertilizer: -1,
		Water:      -1,
		Light:      -1,
		Temp:       -1,
		Humidity:   -1,
		Loc:        -1,
	}
}
