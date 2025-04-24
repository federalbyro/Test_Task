package repository

import "context"

func (r *PostgreDB) SaveRefreshToken(ctx context.Context, rtID, userID, ipAddress, refreshHash string, flagReuse bool) error {

	query := `INSERT INTO tokens ("access_token_id", "user_id", "ip_address", "refresh_hash", "flag_reuse") 
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(ctx, query, rtID, userID, ipAddress, refreshHash, flagReuse)
	if err != nil {
		return err
	}

	return nil
}
