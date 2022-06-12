package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type LoginNasabahRepository interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, nasabah domain.LoginNasabah) (domain.LoginNasabah, error)
	//FindById(ctx context.Context, tx *sql.Tx, idUser int) (domain.LoginNasabah, error)
}
