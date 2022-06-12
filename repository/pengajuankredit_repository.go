package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type PengajuanKreditRepository interface {
	AjuanKredit(ctx context.Context, tx *sql.Tx, ajuankredit domain.PengajuanKredit) domain.PengajuanKredit
}
