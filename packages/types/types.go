package types

type HealthData struct {
	ExportDate   ExportDate    `xml:"ExportDate"`
	Me           Me            `xml:"Me"`
	Records      []Record      `xml:"Record"`
	Correlations []Correlation `xml:"Correlation"`
	Workouts     []Workout     `xml:"Workout"`
}

type ExportDate struct {
	Value string `xml:"value,attr"`
}

type Me struct {
	DateOfBirth                 string `xml:"HKCharacteristicTypeIdentifierDateOfBirth,attr"`
	BiologicalSex               string `xml:"HKCharacteristicTypeIdentifierBiologicalSex,attr"`
	BloodType                   string `xml:"HKCharacteristicTypeIdentifierBloodType,attr"`
	FitzpatrickSkinType         string `xml:"HKCharacteristicTypeIdentifierFitzpatrickSkinType,attr"`
	CardioFitnessMedicationsUse string `xml:"HKCharacteristicTypeIdentifierCardioFitnessMedicationsUse,attr"`
}

type Record struct {
	Type          string `xml:"type,attr"`
	Unit          string `xml:"unit,attr,omitempty"`
	Value         string `xml:"value,attr,omitempty"`
	SourceName    string `xml:"sourceName,attr"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty"`
	Device        string `xml:"device,attr,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty"`
	StartDate     string `xml:"startDate,attr"`
	EndDate       string `xml:"endDate,attr"`
}

type Correlation struct {
	Type          string `xml:"type,attr"`
	SourceName    string `xml:"sourceName,attr"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty"`
	Device        string `xml:"device,attr,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty"`
	StartDate     string `xml:"startDate,attr"`
	EndDate       string `xml:"endDate,attr"`
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

type ActivitySummary struct {
	DateComponents         string `xml:"dateComponents,attr,omitempty"`
	ActiveEnergyBurned     string `xml:"activeEnergyBurned,attr,omitempty"`
	ActiveEnergyBurnedGoal string `xml:"activeEnergyBurnedGoal,attr,omitempty"`
	ActiveEnergyBurnedUnit string `xml:"activeEnergyBurnedUnit,attr,omitempty"`
	AppleMoveTime          string `xml:"appleMoveTime,attr,omitempty"`
	AppleMoveTimeGoal      string `xml:"appleMoveTimeGoal,attr,omitempty"`
	AppleExerciseTime      string `xml:"appleExerciseTime,attr,omitempty"`
	AppleExerciseTimeGoal  string `xml:"appleExerciseTimeGoal,attr,omitempty"`
	AppleStandHours        string `xml:"appleStandHours,attr,omitempty"`
	AppleStandHoursGoal    string `xml:"appleStandHoursGoal,attr,omitempty"`
}

type ClinicalRecord struct {
	Type             string `xml:"type,attr"`
	Identifier       string `xml:"identifier,attr"`
	SourceName       string `xml:"sourceName,attr"`
	SourceURL        string `xml:"sourceURL,attr"`
	FhirVersion      string `xml:"fhirVersion,attr"`
	ReceivedDate     string `xml:"receivedDate,attr"`
	ResourceFilePath string `xml:"resourceFilePath,attr"`
}

type Audiogram struct {
	Type          string `xml:"type,attr"`
	SourceName    string `xml:"sourceName,attr"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty"`
	Device        string `xml:"device,attr,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty"`
	StartDate     string `xml:"startDate,attr"`
	EndDate       string `xml:"endDate,attr"`
}

type SensitivityPoint struct {
	FrequencyValue string `xml:"frequencyValue,attr"`
	FrequencyUnit  string `xml:"frequencyUnit,attr"`
	LeftEarValue   string `xml:"leftEarValue,attr,omitempty"`
	LeftEarUnit    string `xml:"leftEarUnit,attr,omitempty"`
	RightEarValue  string `xml:"rightEarValue,attr,omitempty"`
	RightEarUnit   string `xml:"rightEarUnit,attr,omitempty"`
}

type VisionPrescription struct {
	Type           string     `xml:"type,attr"`
	DateIssued     string     `xml:"dateIssued,attr"`
	ExpirationDate string     `xml:"expirationDate,attr,omitempty"`
	Brand          string     `xml:"brand,attr,omitempty"`
	RightEye       EyeDetails `xml:"RightEye"`
	LeftEye        EyeDetails `xml:"LeftEye"`
	Attachment     Attachment `xml:"Attachment"`
}

type EyeDetails struct {
	Sphere       string `xml:"sphere,attr,omitempty"`
	SphereUnit   string `xml:"sphereUnit,attr,omitempty"`
	Cylinder     string `xml:"cylinder,attr,omitempty"`
	CylinderUnit string `xml:"cylinderUnit,attr,omitempty"`
}

type Attachment struct {
	Identifier string `xml:"identifier,attr,omitempty"`
}
