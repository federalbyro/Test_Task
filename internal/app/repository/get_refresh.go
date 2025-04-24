package repository

import "context"

func (r *PostgreDB) GetRefreshToken(ctx context.Context, rtid string) (string, string, string, bool, error) {
	query := `SELECT "ip_address","flag_reuse","refresh_hash","user_id" FROM tokens WHERE access_token_id = $1`

	var flagReuse bool
	var ipAddress, oldHash, userID string
	err := r.db.QueryRow(ctx, query, rtid).Scan(&ipAddress, &flagReuse, &oldHash, &userID)
	if err != nil {
		return "", "", "", false, err
	}

	return oldHash, ipAddress, userID, flagReuse, nil
}
