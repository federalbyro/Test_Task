package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBRegister interface {
	SaveRefreshToken(ctx context.Context,tokenID, userID, ipAddress, refreshHash string, flag bool) error
	GetRefreshToken(ctx context.Context, oldhash string)(string,string,string,bool,error)
	UpdateRefreshToken(ctx context.Context,oldRtID, newRtid, newIP,newRefreshHash,userID string) error
}

type PostgreDB struct {
	db *pgxpool.Pool
}

func NewPostrgeDB(connect *pgxpool.Pool)*PostgreDB{
	return &PostgreDB{
		db: connect,
	}
}
