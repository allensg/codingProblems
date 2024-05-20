// package warehouse contains the warehouse server code and related structs/utilities for assessing vehicle usage.
package warehouse

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
	"time"
)

// Server implements the warehouse server APIs.
type Server struct {
	vehicles []*Vehicle
}

// AverageSpeeds returns a map from vehicle name to that vehicle's average speed for all vehicles.
func (s *Server) AverageSpeeds() map[string]float64 {
	speeds := make(map[string]float64)
	for _, v := range s.vehicles {
		speeds[v.Name] = v.AverageSpeed()
	}
	return speeds
}

// MostTraveledSince returns a sorted slice of maxResults
// vehicle names corresponding to the vehicles that have traveled the
// most distance since the given timestamp.
func (s *Server) MostTraveledSince(maxResults int, timestamp time.Time) []string {

	// ---------------------------------------------------
	// RUNNING NOTES: i think this might be sorting incorrectly, the results are MBA but should be CAB. with more time would look into this.
	// ---------------------------------------------------

	//so the idea for this is that we'll need to do some pre processing on it.
	// We will need to look at each vehicle and get the sub array of Pings[timestamp:]
	// get the total distance on that sub array and save it to the distance array

	// using a map is kind of risky, but since we're also using float64's
	// the chances of having a duplicate collison is so small that im comfortable with it.
	nameMap, distanceArr := map[float64]string{}, []float64{}

	for _, vehicle := range s.vehicles {
		// get the name and vehicle distances
		distance := vehicle.SubDistance(timestamp)
		nameMap[distance] = vehicle.Name

		distanceArr = append(distanceArr, distance)
	}

	// sort the floats
	sort.Float64s(distanceArr)

	// sort that, and you can iterate over it up to [:maxResults]
	// build the return string array while iterating.
	toReturn := []string{}
	for dIndex, distance := range distanceArr {
		// if we have hit the max results to return, break out of the loop.
		if dIndex == maxResults {
			break
		}
		// since we just populated the map with this array, we shouldn't need to worry about
		// nil collisons.
		name := nameMap[distance]
		toReturn = append(toReturn, name)
	}

	return toReturn
}

// CheckForDamage returns a list of names identifying vehicles that might have been damaged through any number of risky behaviors, including collision with another vehicle and excessive acceleration.
func (s *Server) CheckForDamage() []string {

	// so i don't think im going to have time to solve this but my idea is that we would want a robust
	// set of functions for this. we would need to set up a detector and processor for each heuristic.

	// looking back at the examples here

	// you could look at acceleration and deceleration
	// by setting some kind of threshold in the detector then considering that threshhold as
	// you scan through the pings (IE jumped more than x m/s across a single or range of pings.)

	// check for damage could be handled by a similar threshold and when one is detected, get the timestamp and scan the other vehicles for jumps around that same timestamp.
	// if none are detected it might have been a collision with a wall or object instead of another vehicle.

	// we could make an effort to try to consolidate the checks but this seems
	// like the kind of thing we would want to be robust so i feel like letting it run at the end of the day
	// and run all the detector/processor pairs for each protocol would be more effective.

	// Though another approach could be to embed some of these thresholds and triggering statistics in a struct
	// and add more passive implementations for alerting across total distance, average speed, etc.

	return nil
}

func (s *Server) processPing(vehicleName string, x, y float64, time time.Time) {
	p := &Ping{Position: &Position{X: x, Y: y}, Time: time}
	if len(s.vehicles) == 0 || vehicleName != s.vehicles[len(s.vehicles)-1].Name {
		s.vehicles = append(s.vehicles, &Vehicle{Name: vehicleName})

	}
	s.vehicles[len(s.vehicles)-1].Pings = append(s.vehicles[len(s.vehicles)-1].Pings, p)
}

// NewServerFromFile reads an appropriately-formatted CSV of ping data and returns a new Server initialized with the loaded data.
func NewServerFromFile(filename string) (*Server, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(file)
	s := &Server{}
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		name := record[0]
		x, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		timestamp, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}
		s.processPing(name, float64(x), float64(y), time.Unix(int64(timestamp), 0))
	}
	return s, nil
}
