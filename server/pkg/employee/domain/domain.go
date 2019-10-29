package domain

type Employee struct {
	ID                 uint64 `json:"id"`
	Name               string `json:"name"`
	Salary             uint64 `json:"salary"`
	Position           string `json:"position"`
	Grade              int    `json:"grade"`
	IsTradeUnionMember bool   `json:"is-trade-union-member"`
}
