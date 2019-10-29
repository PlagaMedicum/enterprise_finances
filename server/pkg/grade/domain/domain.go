package domain

type Grade struct {
	ID          uint64 `json:"id"`
	Coefficient map[string]uint64
}
