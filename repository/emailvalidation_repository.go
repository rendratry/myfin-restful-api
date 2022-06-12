package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type EmailValidationRepository interface {
	EmailValidation(ctx context.Context, tx *sql.Tx, validation domain.EmailValidation) (domain.EmailValidation, error)
	FindById(ctx context.Context, tx *sql.Tx, idUser int) (domain.EmailValidation, error)
}
