package domain

import "time"

type Employee struct {
	ID           uint64    `json:"id"`
	Date         time.Time `json:"date"`
	Name         string    `json:"name"`
	Position     string    `json:"position"`
	Grade        uint64    `json:"grade"`
	TUMembership bool      `json:"tu-membership" db:"tu_membership"` // true if employee is trade union member
	Salary       float64   `json:"salary"`
	Accruals     float64   `json:"accruals"`
	Deduction    float64   `json:"deduction"`
}
