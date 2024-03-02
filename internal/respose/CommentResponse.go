package response

type RatingsResponse struct {
	ID           uint
	BorrowerName string  `json:"borrower_name" form:"borrower_name"`
	Message      string  `json:"message" form:"message"`
	Ratings      float32 `json:"ratings" form:"ratings"`
}
