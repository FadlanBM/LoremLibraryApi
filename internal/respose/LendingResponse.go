package response

type LendingResponse struct {
	ID         uint
	DateLast   string `json:"lastdate" form:"lastdate"`
	ReturnDate string `json:"returnDate" form:"returnDate"`
	BorrowDate string `json:"borrowdate" form:"borrowdate"`
	Code       string `json:"code"`
	Status     string `json:"status" form:"status"`
}
