package request

type ListLending struct {
	BookID       uint   `json:"book_id"`
	LendingID    uint   `json:"lending_id"`
	NoInventaris string `json:"no_inventaris"`
}
