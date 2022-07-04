package domain

type TransaksiNasabah struct {
	JenisTransaksi  string
	IdUser          int
	IdTransaksi     int
	Tanggal         string
	Nominal         string
	NominalDiterima string
	Status          string
	Catatan         string
}
