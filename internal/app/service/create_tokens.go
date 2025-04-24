package service

import (
	"context"

	"github.com/google/uuid"
)

func(s *ServiceWorker)GenerateToken(ctx context.Context,userID, ipAddress string)(
	string,
	string,
	string,
	string,
	error, 
){
	refreshAccessID := uuid.New().String()

	accessToken, err := s.tokenGen.CreateAccessToken(userID, refreshAccessID, ipAddress)
	if err != nil {
		return "", "","","", err
	}

	refreshToken,newHash, err := s.tokenGen.GenerateRefreshToken(refreshAccessID)
	if err != nil {
		return "", "","","", err
	}
	return refreshAccessID,accessToken,refreshToken,newHash, nil

}

func (s *ServiceWorker) CreateTokens(ctx context.Context,userID, ipAddress string) (string, string, error) {

	rtID,access,refresh,hash,err := s.GenerateToken(ctx,userID,ipAddress)
	if err!=nil{
		return "","",err
	}
	flagReuse:=false
	err = s.Repository.SaveRefreshToken(ctx, rtID, userID, ipAddress, hash,flagReuse)
	if err!=nil{
		return "","",err
	}
	return access, refresh, err
}
