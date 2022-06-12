package web

type EmailValidationRequest struct {
	IdUser  int    `json:"id"`
	Email   string `json:"email"`
	KodeOtp string `json:"otp"`
	Status  string `json:"status"`
}
