package helper

import (
	"myfin/entity/domain"
	"myfin/entity/web"
)

func ToDataNasabahResponse(nasabah domain.DataNasabah) web.DataNasabahResponse {
	return web.DataNasabahResponse{
		Id_user:      nasabah.IdUser,
		Nik:          nasabah.Nik,
		Nama_lengkap: nasabah.NamaLengkap,
		Alamat:       nasabah.Alamat,
		No_hp:        nasabah.NoHp,
		Tgl_lahir:    nasabah.TglLahir,
		Kota_lahir:   nasabah.KotaLahir,
		KtpFile:      nasabah.KtpFile,
		SwaFile:      nasabah.SwaFile,
	}
}
func ToAkunNasabahResponse(nasabah domain.AkunNasabah) web.AkunNasabahResponse {
	return web.AkunNasabahResponse{
		Id_user: nasabah.IdUser,
		Email:   nasabah.Email,
	}
}

func ToAvaUploadResponse(nasabah domain.Ava) web.AvaUploadResponse {
	return web.AvaUploadResponse{
		Id_user: nasabah.Id,
		Ava:     nasabah.Ava,
	}
}

func ToLoginNasabahResponse(nasabah domain.LoginNasabah) web.LoginNasabahResponse {
	return web.LoginNasabahResponse{
		IdNasabah: nasabah.IdNasabah,
		Email:     nasabah.Email,
		Status:    nasabah.Status,
	}
}

func ToKeamananNasabahResponse(nasabah domain.KeamananNasabah) web.KeamananNasabahResponse {
	return web.KeamananNasabahResponse{
		Id_user:             nasabah.Id,
		Pertanyaan_Keamanan: nasabah.PertanyaanKeamanan,
		Jawaban_Keamanan:    nasabah.JawabanKeamanan,
	}

}
func ToEmailValidationResponse(nasabah domain.EmailValidation) web.EmailValidationResponse {
	return web.EmailValidationResponse{
		Email:   nasabah.Email,
		KodeOtp: nasabah.KodeOtp,
		Status:  nasabah.Status,
	}
}
func ToPinNasabahResponse(nasabah domain.PinNasabah) web.PinNasabahResponse {
	return web.PinNasabahResponse{
		Id:  nasabah.IdUser,
		Pin: nasabah.PinUser,
	}
}

func ToPengajuanKreditResponse(pengajuankredit domain.PengajuanKredit) web.PengajuanKreditResponse {
	return web.PengajuanKreditResponse{
		IdNasabah:       pengajuankredit.IdNasabah,
		Penggunaan:      pengajuankredit.Penggunaan,
		TanggalPengjuan: pengajuankredit.TanggalPengajuan,
		Gaji:            pengajuankredit.Gaji,
		GajiTambahan:    pengajuankredit.GajiTambahan,
		BesarPengajuan:  pengajuankredit.BesarPengajuan,
		Tenor:           pengajuankredit.Tenor,
		Score:           pengajuankredit.Score,
		Status:          pengajuankredit.Status,
	}
}

func ToPenarikanSaldoResponse(penarikansaldo domain.PenarikanSaldo) web.PenarikanSaldoResponse {
	return web.PenarikanSaldoResponse{
		IdNasabah:       penarikansaldo.IdNasabah,
		JumlahPenarikan: penarikansaldo.JumlahPenarikan,
		Bank:            penarikansaldo.Bank,
		NamaPemilik:     penarikansaldo.NamaPemilik,
		NoRek:           penarikansaldo.NoRek,
		Status:          penarikansaldo.Status,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToGetSaldoResponse(saldo domain.Saldo) web.GetSaldoResponse {
	return web.GetSaldoResponse{
		IdNasabah: saldo.IdNasabah,
		Saldo:     saldo.Saldo,
	}
}

func ToGetNikResponse(nik domain.Nik) web.GetNikResponse {
	return web.GetNikResponse{
		Nik:          nik.Nik,
		Nama:         nik.Nama,
		Alamat:       nik.Alamat,
		TanggalLahir: nik.TanggalLahir,
		KotaLahir:    nik.KotaLahir,
	}
}

func ToTransaksiResponse(transaksi domain.TransaksiNasabah) web.TransaksiResponse {
	//if transaksi.NominalDiterima == nil{
	//	transaksi.NominalDiterima = string(0)
	//}else{
	//	transaksi.NominalDiterima = transaksi.NominalDiterima
	//}
	return web.TransaksiResponse{

		IdTransaksi:     transaksi.IdTransaksi,
		Tanggal:         transaksi.Tanggal,
		Nominal:         transaksi.Nominal,
		NominalDiterima: transaksi.NominalDiterima,
		Status:          transaksi.Status,
		Catatan:         transaksi.Catatan,
	}
}

func ToGetProfileResponse(id_user domain.DataNasabah) web.GetProfileResponse {
	return web.GetProfileResponse{
		Nama:   id_user.NamaLengkap,
		Email:  id_user.Email,
		Ava:    id_user.Ava,
		NoHp:   id_user.NoHp,
		Nik:    id_user.Nik,
		Alamat: id_user.Alamat,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
