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
	currentState     ElevatorState
	currentFloor     int
	currentDirection Direction
	UpQueue          map[int]bool
	DownQueue        map[int]bool
	capacity         int
	mu               sync.RWMutex
}

func NewElevator() *Elevator {
	return &Elevator{ID: uuid.NewString(), currentState: 0, currentFloor: 0, currentDirection: 0, UpQueue: make(map[int]bool), DownQueue: make(map[int]bool)}
}

func (e *Elevator) Run() {
	for {
		time.Sleep(1*time.Second)
		e.mu.Lock()
		switch e.currentDirection {
			case Up : e.currentFloor++
			case Down : e.currentFloor--
		}
		e.ProcessSteps()
		e.mu.Unlock()
	}
}

func(e *Elevator)ProcessSteps(){
	if e.UpQueue[e.currentFloor] {
		delete(e.UpQueue, e.currentFloor)
	} else if e.DownQueue[e.currentFloor] {
		delete(e.DownQueue, e.currentFloor)
	}

	if e.currentDirection==Up && len(e.UpQueue)==0 {
		if len(e.DownQueue)>0 {
			e.currentDirection=Down
			} else {
			e.currentDirection=Idle

		}
	} else if e.currentDirection==Down && len(e.DownQueue)==0 {
		if len(e.UpQueue)>0 {
			e.currentDirection=Up
			} else {
			e.currentDirection=Idle

		}
	}

}

func (e *Elevator) AddStop(floor int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if floor > e.currentFloor {
		e.UpQueue[floor]=true
	} else if floor < e.currentFloor {
		e.DownQueue[floor]=true
	} 

	if e.currentDirection == Idle {
		if floor > e.currentFloor {
			e.currentDirection = Up
		} else {
			e.currentDirection = Down
		}
	}
}
