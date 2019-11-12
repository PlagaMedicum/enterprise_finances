package domain

import "math/big"

type Employee struct {
	ID           big.Int
	Name         string
	Position     string
	Grade        big.Int
	TUMembership bool `json:"is-trade-union-member" db:"tu_membership"` // true if employee is trade union member
}
