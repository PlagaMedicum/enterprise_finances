package handlers

type Usecases interface {
	AddEmployee() error
	EditEmployee() error
	DeleteEmployee() error
	GetEmployeeList() error
}
