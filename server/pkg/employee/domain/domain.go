package domain

type Employee struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Grade        uint64 `json:"grade"`
	TUMembership bool   `json:"tu-membership" db:"tu_membership"` // true if employee is trade union member
	Salary       map[string]uint64
}
