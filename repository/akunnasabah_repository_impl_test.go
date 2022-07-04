package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"myfin/helper"
	"testing"
)

func TestAkunNasabahInsert(t *testing.T) {
	//akunnasabahRepository := NewAkunNasabahRepository(app.GetConnection())
	//
	//ctx := context.Background()
	//akunnasabah := domain.DataNasabah{
	//	Email:    "rendratrikusuma1@gmail.com",
	//	Kode_otp: 4567,
	//}
	//result, err := akunnasabahRepository.Insert(ctx, akunnasabah)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)
}

func TestSendEmail(t *testing.T) {
	helper.SendEmail("mencoba", "rendratrykusuma@gmail.com", "hai")
}
