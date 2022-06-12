package web

type PengajuanKreditRequest struct {
	IdNasabah      int    `validate:"required" json:"id"`
	Penggunaan     string `validate:"required" json:"penggunaan"`
	BesarPengajuan int    `validate:"required" json:"besar_pengajuan"`
	Tenor          string `validate:"required" json:"tenor"`
	Score          int    `validate:"required" json:"score"`
	//Status         string `validate:"required" json:"status"`
}
