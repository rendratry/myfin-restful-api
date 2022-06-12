package web

type LoginNasabahRequest struct {
	Email string `json:"email"`
	Pin   string `json:"pin"`
}
