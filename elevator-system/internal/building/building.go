package building

import (
	"elevator-system/internal/dispatcher"
	"elevator-system/internal/elevator"
)

type Building struct {
	Elevators  []*elevator.Elevator
	Dispatcher *dispatcher.Dispatcher
}

func NewBuilding(elevators []*elevator.Elevator, disp *dispatcher.Dispatcher) *Building {
	if elevators == nil {
		elevators = []*elevator.Elevator{}
	}
	if disp == nil {
		disp = dispatcher.NewDispatcher()
	}
	return &Building{Elevators: elevators, Dispatcher: disp}
}
