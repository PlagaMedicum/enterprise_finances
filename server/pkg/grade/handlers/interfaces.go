package handlers

type Usecases interface {
	AddInfo() error
	EditInfo() error
	DeleteInfo() error
	GetGradeList() error
}
