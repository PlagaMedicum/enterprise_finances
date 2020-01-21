package domain

import "time"

type Grade struct {
	ID          uint64    `json:"id"`
	Num         uint64    `json:"num"`
	Date        time.Time `json:"date"`
	Coefficient int       `json:"coeff" db:"coeff" bson:"coeff"`
}
