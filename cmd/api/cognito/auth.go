package cognito

type UserCredentials struct {
	Name     string
	Password string
}

type UserSession struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
	ExpiresIn    int64
}

type AccessToken string

func (t AccessToken) Value() *string {
	v := string(t)
	return &v
}

type RefreshToken string

func (t RefreshToken) Value() *string {
	v := string(t)
	return &v
}
