package web

type PenarikanSaldoRequest struct {
	IdNasabah       int    `validate:"required" json:"id"`
	JumlahPenarikan int    `validate:"required" json:"jml_penarikan"`
	Bank            string `validate:"required" json:"bank"`
	NoRek           string `validate:"required" json:"norek"`
	NamaPemilik     string `validate:"required" json:"nama_pemilik"`
}
