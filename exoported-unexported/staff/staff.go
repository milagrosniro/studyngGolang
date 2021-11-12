package staff

//TODO LO QUE EMPIEZA CON MAYUSCULA ES EXPORTABLE

var overPaidLimit = 75000 //no exportable
var underpaidLimit = 5000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	Fulltime  bool
}

type Office struct {
	AllStaf []Employee
}

//funcion que recibe un puntero a office y devuelve un array de employee
//Ahora Office tiene una funcion lladama All()
func (e *Office) All() []Employee {
	return e.AllStaf
}

func (e *Office) OverPaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaf {
		if x.Salary > overPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) UnderPaid() []Employee {
	var underpaid []Employee

	for _, x := range e.AllStaf {
		if x.Salary < underpaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}
