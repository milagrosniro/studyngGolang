package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "jonh", LastName: "DIAz", Salary: 50000, Fulltime: true},
	{FirstName: "sally", LastName: "jones", Salary: 60000, Fulltime: true},
	{FirstName: "mark", LastName: "smith", Salary: 30000, Fulltime: true},
	{FirstName: "jonh", LastName: "smith", Salary: 85000, Fulltime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaf: employees,
	}
	//	log.Println(myStaff.All())
	log.Println(myStaff.OverPaid())

}
