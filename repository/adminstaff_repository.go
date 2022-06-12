package repository

import (
	"context"
	"myfin/entity/domain"
)

type AdminstaffRepository interface {
	Insert(ctx context.Context, adminstaff domain.AdminStaff) (domain.AdminStaff, error)
	FindById(ctx context.Context, idadminstaff domain.AdminStaff) (AdminstaffRepository, error)
}
