package domain

import "math/big"

type Grade struct {
	ID          big.Int `json:"id"`
	Date        string  `json:"date"`
	Coefficient int     `json:"coefficient"`
}
