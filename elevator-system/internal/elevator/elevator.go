package elevator

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type ElevatorState int
type Direction int

const (
	STOPPED ElevatorState = iota
	MOVING
)

const (
	Idle Direction = iota
	Up
	Down
)

type Elevator struct {
	ID               string
	CurrentState     ElevatorState
	CurrentFloor     int
	CurrentDirection Direction
	UpQueue          map[int]bool
	DownQueue        map[int]bool
	capacity         int
	Mu               sync.RWMutex
}

func NewElevator() *Elevator {
	return &Elevator{ID: uuid.NewString(), CurrentState: 0, CurrentFloor: 0, CurrentDirection: 0, UpQueue: make(map[int]bool), DownQueue: make(map[int]bool)}
}

func (e *Elevator) Run() {
	for {
		time.Sleep(1*time.Second)
		e.Mu.Lock()
		switch e.CurrentDirection {
			case Up : e.CurrentFloor++
			case Down : e.CurrentFloor--
		}
		e.ProcessSteps()
		e.Mu.Unlock()
	}
}

func(e *Elevator)ProcessSteps(){
	if e.UpQueue[e.CurrentFloor] {
		delete(e.UpQueue, e.CurrentFloor)
	} else if e.DownQueue[e.CurrentFloor] {
		delete(e.DownQueue, e.CurrentFloor)
	}

	if e.CurrentDirection==Up && len(e.UpQueue)==0 {
		if len(e.DownQueue)>0 {
			e.CurrentDirection=Down
			} else {
			e.CurrentDirection=Idle
		}
	} else if e.CurrentDirection==Down && len(e.DownQueue)==0 {
		if len(e.UpQueue)>0 {
			e.CurrentDirection=Up
			} else {
			e.CurrentDirection=Idle

		}
	}

}

func (e *Elevator) AddStop(floor int) {
	e.Mu.Lock()
	defer e.Mu.Unlock()
if floor == e.CurrentFloor {
    return
}

	if floor > e.CurrentFloor {
		e.UpQueue[floor]=true
	} else if floor < e.CurrentFloor {
		e.DownQueue[floor]=true
	} 

	if e.CurrentDirection == Idle {
		if floor > e.CurrentFloor {
			e.CurrentDirection = Up
		} else {
			e.CurrentDirection = Down
		}
	}
}
