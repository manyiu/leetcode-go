package main

// ParkingSystem is the given struct
type ParkingSystem struct {
	big int
	medium int
	small int
}

// Constructor creates car park
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

// AddCar parks an extra car
func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.big > 0 {
			p.big--
			return true
		}
	case 2:
		if p.medium > 0 {
			p.medium--
			return true
		}
	case 3:
		if p.small > 0 {
			p.small--
			return true
		}
	default:
		return false
	}

	return false
}

func main() {}