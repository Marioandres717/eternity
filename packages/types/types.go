package types

type HealthData struct {
	Workouts []Workout `xml:"Workout"`
}

type Workout struct {
	WorkoutActivityType   string `xml:"workoutActivityType,attr"`
	Duration              string `xml:"duration,attr,omitempty"`
	DurationUnit          string `xml:"durationUnit,attr,omitempty"`
	TotalDistance         string `xml:"totalDistance,attr,omitempty"`
	TotalDistanceUnit     string `xml:"totalDistanceUnit,attr,omitempty"`
	TotalEnergyBurned     string `xml:"totalEnergyBurned,attr,omitempty"`
	TotalEnergyBurnedUnit string `xml:"totalEnergyBurnedUnit,attr,omitempty"`
	SourceName            string `xml:"sourceName,attr"`
	SourceVersion         string `xml:"sourceVersion,attr,omitempty"`
	Device                string `xml:"device,attr,omitempty"`
	CreationDate          string `xml:"creationDate,attr,omitempty"`
	StartDate             string `xml:"startDate,attr"`
	EndDate               string `xml:"endDate,attr"`
}
