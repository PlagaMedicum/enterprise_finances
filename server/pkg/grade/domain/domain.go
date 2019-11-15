package domain

type Grade struct {
	ID          uint64 `json:"id"`
	Date        string `json:"date"`
	Coefficient int    `json:"coefficient" db:"coeff"`
}
