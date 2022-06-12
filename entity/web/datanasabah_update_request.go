package web

type DatanasabahUpdateRequest struct {
	Id_user      int    `validate:"required"`
	Nik          string `validate:"required" json:"nik"`
	Nama_lengkap string `validate:"required,max=200,min=1" json:"nama"`
	Alamat       string `validate:"required" json:"alamat"`
	No_hp        string `validate:"required" json:"no_hp"`
	Tgl_lahir    string `validate:"required" json:"lahir"`
	Kota_lahir   string `validate:"required" json:"kota"`
	KtpFile      string `validate:"required" json:"ktp"`
	SwaFile      string `validate:"required" json:"swa"`
}

type DatanasabahUpdatePinRequest struct {
	Id_user int    `validate:"required"`
	Pin     string `validate:"required" json:"pin"`
}

type AvaUploadRequest struct {
	Id_user int    `validate:"required"`
	Ava     string `validate:"required" json:"ava"`
}

type ProfilenasabahRequest struct {
	Id_user      int    `validate:"required"`
	Email        string `validate:"required" json:"email"`
	Nik          string `validate:"required" json:"nik"`
	Nama_lengkap string `validate:"required,max=200,min=1" json:"nama"`
	Alamat       string `validate:"required" json:"alamat"`
	No_hp        string `validate:"required" json:"no_hp"`
}

type MinSaldoUpdateRequest struct {
	IdNasabah int `validate:"required"`
	Saldo     int `validate:"required" json:"saldo"`
}

type DatanasabahUpdateKeamananRequest struct {
	Id_user             int    `validate:"required"`
	Pertanyaan_Keamanan string `validate:"required" json:"pertanyaan"`
	Jawaban_Keamanan    string `validate:"required" json:"jawaban"`
}
