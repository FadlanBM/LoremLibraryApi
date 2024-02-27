package request

type CollectionRequest struct {
	BorrowersID uint `json:"borrowers_id"`
	BooksID     uint `json:"books_id"`
}
