package request

type AuthRequest struct {
	GoogleID string `json:"google_id"`
	Email string `json:"email"`
}
