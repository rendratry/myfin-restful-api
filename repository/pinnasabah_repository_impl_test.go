package repository

import (
	_ "github.com/go-sql-driver/mysql"
	_ "golang.org/x/crypto/bcrypt"
	"testing"
)

func TestInsertPinNasabah(t *testing.T) {
	//pinnasabahrepository := NewPinNasabahRepository(app.GetConnection())
	//ctx := context.Background()
	//
	//pass := []byte("578490")
	//
	//// Hashing the password
	//hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	//if err != nil {
	//	panic(err)
	//}
	//pinnasabah := domain.DataNasabah{
	//	Pin_user: string(hash),
	//	Id:  1,
	//}
	//result, err := pinnasabahrepository.Insert(ctx, pinnasabah)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(result)
}
