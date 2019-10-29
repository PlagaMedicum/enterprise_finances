package usecases

type Repository interface {
	AddEmployee() error
	UpdateEmployee() error
	DeleteEmployee() error
	GetEmployeeList() error
}
