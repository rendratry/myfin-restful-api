package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type adminStaffRepositoryImpl struct {
	DB *sql.DB
}

func (repository adminStaffRepositoryImpl) Insert(ctx context.Context, adminstaff domain.AdminStaff) (domain.AdminStaff, error) {
	script := "INSERT INTO tb_adminstaff(email_adminstaff, password_adminstaff, nama_adminstaff, role) VALUES (?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, adminstaff.Email_adminstaff, adminstaff.Password_adminstaff, adminstaff.Nama_adminstaff, adminstaff.Role)
	if err != nil {
		return adminstaff, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return adminstaff, err
	}
	adminstaff.Id_adminstaff = int32(id)
	return adminstaff, nil
}

func (repository adminStaffRepositoryImpl) FindById(ctx context.Context, idadminstaff domain.AdminStaff) (AdminstaffRepository, error) {
	//TODO implement me
	panic("implement me")
}
