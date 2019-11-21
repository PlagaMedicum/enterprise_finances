package domain

type Employee struct {
	ID           uint64
	Name         string
	Position     string
	Grade        uint64
	TUMembership bool `json:"tu-membership" db:"tu_membership"` // true if employee is trade union member
	Salary       map[string]uint64
}
