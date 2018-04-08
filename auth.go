package tchat

type Account struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Register(acc *Account) error
	Auth(acc *Account) error
	Refresh(acc *Account) error
}
