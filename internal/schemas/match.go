package schemas

type Match struct {
	Id    Guid   `json:"id"`
	Board Board  `json:"board"`
	Moves []Move `json:"moves"`
}
