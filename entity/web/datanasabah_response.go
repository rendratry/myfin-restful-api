package web

type DataNasabahResponse struct {
	Id_user             int    `json:"id"`
	Email               string `json:"email"`
	Pertanyaan_Keamanan string `json:"pertanyaan"`
	Jawaban_Keamanan    string `json:"jawaban"`
	Nik                 string `json:"nik"`
	Nama_lengkap        string `json:"nama"`
	Alamat              string `json:"alamat"`
	No_hp               string `json:"no_hp"`
	Tgl_lahir           string `json:"lahir"`
	Kota_lahir          string `json:"kota"`
	KtpFile             string `json:"ktp"`
	SwaFile             string `json:"swa"`
}
type AkunNasabahResponse struct {
	Id_user int    `json:"id"`
	Email   string `json:"email"`
}

type AvaUploadResponse struct {
	Id_user int    `json:"id"`
	Ava     string `json:"ava"`
}

type KeamananNasabahResponse struct {
	Id_user             int    `json:"id"`
	Pertanyaan_Keamanan string `json:"pertanyaan"`
	Jawaban_Keamanan    string `json:"jawaban"`
}
type PinNasabahResponse struct {
	Id  int    `json:"id"`
	Pin string `json:"pin"`
}

type LoginNasabahResponse struct {
	IdNasabah int    `json:"id"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

type EmailValidationResponse struct {
	Email   string `json:"email"`
	KodeOtp string `json:"otp"`
	Status  string `json:"status"`
}

type EmailCheckResponse struct {
}

type PenarikanSaldoResponse struct {
	IdNasabah       int    `json:"id"`
	JumlahPenarikan int    `json:"jml_penarikan"`
	Bank            string `json:"bank"`
	NoRek           string `json:"norek"`
	NamaPemilik     string `json:"nama_pemilik"`
	Status          string `json:"status"`
}

type GetSaldoResponse struct {
	IdNasabah int `json:"id"`
	Saldo     int `json:"saldo"`
}

type GetNikResponse struct {
	Nik          int    `json:"nik"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tgl_lahir"`
	KotaLahir    string `json:"kota"`
}

type GetProfileResponse struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Ava    string `json:"ava"`
	NoHp   string `json:"no_hp"`
	Nik    string `json:"nik"`
	Alamat string `json:"alamat"`
}

type PengajuanKreditResponse struct {
	IdNasabah       int    `json:"id"`
	Penggunaan      string `json:"penggunaan"`
	BesarPengajuan  int    `json:"besar_pengajuan"`
	Tenor           string `json:"tenor"`
	Score           int    `json:"score"`
	TanggalPengjuan string `json:"tanggal_pengjuan"`
	Status          string `json:"status"`
}
