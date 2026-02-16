package dispatcher

import (
	"elevator-system/internal/elevator"
)

type Dispatcher struct {
	Elevators []*elevator.Elevator
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{Elevators: []*elevator.Elevator{}}

}

func (d *Dispatcher) AddElevator(e *elevator.Elevator) {
	d.Elevators = append(d.Elevators, e)
}

func (d *Dispatcher) FindBestElevator(floor int, direction *elevator.Direction) *elevator.Elevator {
	var BestElevator *elevator.Elevator
	var minScore int = -1
	for _, e := range d.Elevators {
		e.Mu.Lock()
		score := ScoreThisElevator(e, floor, direction)
		e.Mu.Unlock()

		if score < minScore {
			minScore = score
			BestElevator = e
		}
	}
	return BestElevator
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func ScoreThisElevator(e *elevator.Elevator, floor int, direction *elevator.Direction) int {
	score := abs(floor - e.CurrentFloor)

	if (e.CurrentDirection == elevator.Up && *direction == elevator.Up) || (e.CurrentDirection == elevator.Down && *direction == elevator.Down) {
		score -= 2
	}
	if e.CurrentDirection == elevator.Idle {
		score -= 5
	}
	return score

}
