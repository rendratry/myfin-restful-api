package web

type DatanasabahCreateRequest struct {
	Email    string `validate:"email" json:"email"`
	Kode_otp string `validate:"required" json:"otp"`
}
