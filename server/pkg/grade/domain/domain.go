package domain

type Grade struct {
	ID          uint64 `json:"id"`
	Num         uint64 `json:"num"`
	Date        string `json:"date"`
	Coefficient int    `json:"coeff" db:"coeff"`
}
