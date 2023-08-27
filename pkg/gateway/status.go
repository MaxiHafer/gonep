package gateway

type Status struct {
	Id                  string
	CurrentWatts        int
	TodayWattHours      int
	TotalWattHours      int
	KilogramsOfCO2Saved int
	Status              string
}
