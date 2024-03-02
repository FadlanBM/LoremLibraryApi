package request

type CommentRequest struct {
	BooksID  uint    `json:"buku_id"`
	Rating   float32 `json:"ratings"`
	Messages string  `json:"messages"`
}
