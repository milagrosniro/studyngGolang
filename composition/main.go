package main

import "fmt"

type Vehicle struct {
	NumberOdWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passenger", v.NumberOfPassengers)
	fmt.Println("Number of wheels", v.NumberOdWheels)

}

func (c Car) show() {
	fmt.Println("Make", c.Make)
	fmt.Println("Model", c.Model)
	fmt.Println("Year", c.Year)
	fmt.Println("Electric", c.IsElectric)
	fmt.Println("Hybrid", c.IsHybrid)
	c.Vehicle.showDetails()

}

func main() {
	suv := Vehicle{
		NumberOdWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "volvo",
		Model:      "XC90",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()

	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "ModelX",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	teslaModelX.Vehicle.NumberOfPassengers = 7
	teslaModelX.show()
}
