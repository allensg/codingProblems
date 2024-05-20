package warehouse

import (
	"fmt"
	"time"
)

// Ping represents a vehicle's position at a given timestamp.
type Ping struct {
	Position *Position
	Time     time.Time
}

// String returns the string representation of a Ping.
func (p *Ping) String() string {
	return fmt.Sprintf("%s @ %d", p.Position, p.Time.Unix())
}

// SecondsBetween determines the number of seconds between two given Pings. The result is positive if ping1 is earlier than ping2.
func SecondsBetween(ping1, ping2 *Ping) float64 {
	return ping2.Time.Sub(ping1.Time).Seconds()
}
