package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestPertanyaanKeamananInsert(t *testing.T) {
	//pertanyaankeamananRepository := NewPertanyaanKeamananRepository(app.GetConnection())
	//ctx := context.Background()
	//
	//pertanyaanKeamanan := domain.DataNasabah{
	//	IdUser:              1,
	//	Pertanyaan_keamanan: "Siapa Nama Gurumu?",
	//	Jawaban_keamanan:    "Juari",
	//}
	//result, err := pertanyaankeamananRepository.Insert(ctx, pertanyaanKeamanan)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)
}
