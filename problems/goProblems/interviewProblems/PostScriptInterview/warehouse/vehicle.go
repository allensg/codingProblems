package warehouse

import (
	"fmt"
	"time"
)

// Vehicle represents a named vehicle with a sequence of pings.
type Vehicle struct {
	Name  string
	Pings []*Ping
}

// TotalDistance returns the total distance traveled by the vehicle.
func (v *Vehicle) TotalDistance() float64 {
	return TotalDistance(v.Pings)
}

// AverageSpeed returns the average speed of the vehicle.
func (v *Vehicle) AverageSpeed() float64 {
	// We will need the (total distance / total time)
	averageSpeed := 0.0

	// Under the assumption 'the input list is sorted with the earliest ping at the start'
	// total time should be just the lastPing.time - firstPing.time
	pings := v.Pings

	// if we only have one record or 0 records we can't compute the average
	if len(pings) <= 1 {
		return 0
	}

	firstPing, lastPing := pings[0], pings[len(pings)-1]
	totalTime := SecondsBetween(firstPing, lastPing)

	// Total distance from the masterfully executed TotalDistance()
	totalDistance := TotalDistance(v.Pings)

	averageSpeed = totalDistance / totalTime

	return averageSpeed
}

// String returns a string representation of the vehicle.
func (v *Vehicle) String() string {
	return fmt.Sprintf("%s: %v", v.Name, v.Pings)
}

// TotalDistance returns the total distance covered by the specified pings.
func TotalDistance(pings []*Ping) float64 {
	// Iterate over the pings, use the built in distance function to determine the
	// distance, total it up and return
	totalDistance := 0.0
	for i := 1; i < len(pings); i++ {
		previousPing := pings[i-1]
		currentPing := pings[i]
		traveled := Distance(previousPing.Position, currentPing.Position)
		totalDistance = totalDistance + traveled
	}

	return totalDistance
}

// SubDistance returns the distance traveled since a given time by the vehicle.
func (v *Vehicle) SubDistance(timestamp time.Time) float64 {
	return SubDistance(v.Pings, timestamp)
}

// SubDistance returns the distance covered by the specified pings over a given time range.
func SubDistance(pings []*Ping, timestamp time.Time) float64 {
	// Iterate over the pings, use the built in distance function to determine the
	// distance, total it up and return
	subDistance := 0.0
	for i := 1; i < len(pings); i++ {

		previousPing := pings[i-1]
		currentPing := pings[i]
		if timestamp.Compare(currentPing.Time) > 0 {
			// if we have not yet reached the time boundary, skip
			continue
		}
		traveled := Distance(previousPing.Position, currentPing.Position)
		subDistance = subDistance + traveled
	}

	return subDistance
}
