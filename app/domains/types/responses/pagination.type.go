package responses

type Pagination struct {
	PreviousPage uint   `json:"previous_page"`
	CurrentPage  uint   `json:"current_page"`
	NextPage     uint   `json:"next_page"`
	MaxPage      int64  `json:"max_page"`
	Query        string `json:"query"`
}
