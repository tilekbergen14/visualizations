package models

type Relation struct {
	Id             string
	DatesLocations []DatesLocations
}

type DatesLocations struct {
	Location string
	Dates    []string
}
