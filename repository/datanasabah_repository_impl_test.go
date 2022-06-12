package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDataNasabahInsert(t *testing.T) {
	//datanasabahRepository := NewDataNasabahRepository(app.GetConnection())
	//
	//ctx := context.Background()
	//datanasabah := domain.DataNasabah{
	//	Nik:         "3519110806000002",
	//	NamaLengkap: "Iwan Pambudi",
	//	NoHp:        "6281776885994",
	//	TglLahir:    "2000-07-09",
	//	KotaLahir:   "Madiun",
	//	IdUser:      1,
	//}
	//result, err := datanasabahRepository.insert(ctx, datanasabah)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)
}
