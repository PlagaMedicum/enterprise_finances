package usecases

type Repository interface {
	AddInfo() error
	EditInfo() error
	DeleteInfo() error
	GetGradeList() error
}
