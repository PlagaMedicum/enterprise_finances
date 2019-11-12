package domain

import "math/big"

type Employee struct {
	ID           big.Int `json:"id"`
	Name         string  `json:"name"`
	Position     string  `json:"position"`
	Grade        big.Int `json:"grade"`
	TUMembership bool    `json:"is-trade-union-member"` // true if employee is trade union member
}
