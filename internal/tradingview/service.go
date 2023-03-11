package tradingview


type Service interface {
	InviteUser(scriptID, username string) error
	RemoveUser(scriptID, username string) error
}