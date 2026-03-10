package internal

type ParkingLot struct {
	levels         []*level
	vehicleIndex   map[string]*Spot
	availableSpots map[VehicleType][]*Spot
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{levels: make([]*level, 0), vehicleIndex: make(map[string]*Spot), availableSpots: make(map[VehicleType][]*Spot)}
}


func(p *ParkingLot) AddLevel(newLevel *level){
	p.levels=append(p.levels, newLevel)
}

