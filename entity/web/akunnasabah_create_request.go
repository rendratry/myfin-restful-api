package web

type AkunNasabahCreateRequest struct {
	Email string `validate:"email" json:"email"`
}
