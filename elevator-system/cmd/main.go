package main

import (
	"elevator-system/internal/building"
	"elevator-system/internal/dispatcher"
	"elevator-system/internal/elevator"
	"fmt"
	"time"
)

func main() {

	// Create elevators
	e1 := elevator.NewElevator()
	e2 := elevator.NewElevator()
	e3 := elevator.NewElevator()

	// Start them as goroutines
	go e1.Run()
	go e2.Run()
	go e3.Run()

	// Create dispatcher
	disp := dispatcher.NewDispatcher()
	disp.AddElevator(e1)
	disp.AddElevator(e2)
	disp.AddElevator(e3)

	// Create building
	bld := building.NewBuilding([]*elevator.Elevator{e1, e2, e3}, disp)

	// Give system time to boot
	time.Sleep(1 * time.Second)

	// Simulate external request
	requestedDirection := elevator.Up
	best := bld.Dispatcher.FindBestElevator(5, &requestedDirection)

	if best != nil {
		fmt.Println("Assigned elevator:", best.ID)
		best.AddStop(5)
	} else {
		fmt.Println("No elevator available")
	}

	// Simulate another request
	time.Sleep(2 * time.Second)

	requestedDirection = elevator.Down
	best = bld.Dispatcher.FindBestElevator(2, &requestedDirection)

	if best != nil {
		fmt.Println("Assigned elevator:", best.ID)
		best.AddStop(2)
	}

	// Keep program alive to observe movement
	for {
		time.Sleep(3 * time.Second)

		e1.Mu.RLock()
		fmt.Println("E1 -> Floor:", e1.CurrentFloor, "Direction:", e1.CurrentDirection)
		e1.Mu.RUnlock()

		e2.Mu.RLock()
		fmt.Println("E2 -> Floor:", e2.CurrentFloor, "Direction:", e2.CurrentDirection)
		e2.Mu.RUnlock()

		e3.Mu.RLock()
		fmt.Println("E3 -> Floor:", e3.CurrentFloor, "Direction:", e3.CurrentDirection)
		e3.Mu.RUnlock()

		fmt.Println("-----")
	}
}
