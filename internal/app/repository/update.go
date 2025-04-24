package repository

import (
	"context"
	"errors"
)

func (r *PostgreDB) UpdateRefreshToken(ctx context.Context, oldRtid, newRtID, newIP, newRefreshHash, userID string) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	query := `
		UPDATE tokens
		SET flag_reuse=true WHERE access_token_id = $1 AND flag_reuse=false
	`

	result, err := tx.Exec(ctx, query, oldRtid)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return errors.New("row already exist")
	}

	secondQuery :=
		`INSERT INTO tokens ("access_token_id", "user_id", "ip_address", "refresh_hash", "flag_reuse") 
	VALUES ($1, $2, $3, $4, false)`
	result, err = tx.Exec(ctx, secondQuery, newRtID, userID, newIP, newRefreshHash)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
