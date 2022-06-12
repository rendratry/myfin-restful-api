package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type AvaUploadRepository interface {
	AvaUpload(ctx context.Context, tx *sql.Tx, ava domain.Ava) domain.Ava
	FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.Ava, error)
}
