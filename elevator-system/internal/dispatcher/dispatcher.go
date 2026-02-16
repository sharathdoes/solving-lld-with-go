package dispatcher

import (
	"elevator-system/internal/elevator"
)
type Dispatcher struct {
	Elevators []*elevator.Elevator
}


func NewDispatcher() *Dispatcher {
	return &Dispatcher{Elevators: []*elevator.Elevator{}};

}

func (d *Dispatcher) AddElevator( e *elevator.Elevator){
	d.Elevators = append(d.Elevators, e)
}

func(d *Dispatcher) FindBestElevator(floor int, direction *elevator.Direction){
	var BestElevator *elevator.Elevator
	var minScore int=-1
	for _,e:=range d.Elevators {
		minScore=min(minScore, ScoreThisElevator(e,floor, direction ))
	}
}

func abs( x int) int{
	
}
func ScoreThisElevator(e *elevator.Elevator, floor int, direction *elevator.Direction) int{
	score:=abs()
}