package repository

type DBRegister interface {
	SaveRefreshToken(tokenID, userID, ipAddress, refreshHash string) error
	GetRefreshToken()
	CheckTokenByIp()
}
