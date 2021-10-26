package cognito

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/pkg/errors"
)

type userOperations interface {
	Register(credentials UserCredentials) error

	Login(credentials UserCredentials) (UserSession, error)
	Refresh(refreshToken RefreshToken) (UserSession, error)

	ForgotPassword(username string) error
	ConfirmForgotPassword(credentials UserCredentials, code string) (UserSession, error)

	DeleteUser(accessToken AccessToken) error
}

func (m *userManager) Register(credentials UserCredentials) error {
	user := &cognitoidentityprovider.SignUpInput{
		ClientId: m.appClientID,
		Password: aws.String(credentials.Password),
		Username: aws.String(credentials.Name),
	}

	_, err := m.cognitoClient.SignUp(user)
	return err
}

func (m *userManager) Login(credentials UserCredentials) (UserSession, error) {
	res, err := m.cognitoClient.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(credentials.Name),
			"PASSWORD": aws.String(credentials.Password),
		},
		ClientId: m.appClientID,
	})
	if err != nil {
		return UserSession{}, err
	}

	return m.toUserSession(res.AuthenticationResult, nil)
}

func (m *userManager) Refresh(refreshToken RefreshToken) (UserSession, error) {
	res, err := m.cognitoClient.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": refreshToken.Value(),
		},
		ClientId: m.appClientID,
	})
	if err != nil {
		return UserSession{}, err
	}

	return m.toUserSession(res.AuthenticationResult, &refreshToken)
}

func (m *userManager) ForgotPassword(username string) error {
	_, err := m.cognitoClient.ForgotPassword(&cognitoidentityprovider.ForgotPasswordInput{
		ClientId: m.appClientID,
		Username: &username,
	})
	return err
}

func (m *userManager) ConfirmForgotPassword(credentials UserCredentials, code string) (UserSession, error) {
	_, err := m.cognitoClient.ConfirmForgotPassword(&cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         m.appClientID,
		ConfirmationCode: &code,
		Password:         &credentials.Password,
		Username:         &credentials.Name,
	})
	if err != nil {
		return UserSession{}, err
	}

	return m.Login(credentials)
}

func (m *userManager) DeleteUser(accessToken AccessToken) error {
	_, err := m.cognitoClient.DeleteUser(&cognitoidentityprovider.DeleteUserInput{
		AccessToken: accessToken.Value(),
	})
	return err
}

func (m *userManager) extractRefreshToken(authResult *cognitoidentityprovider.AuthenticationResultType, refreshToken *RefreshToken) *RefreshToken {
	if authResult == nil || authResult.RefreshToken == nil {
		return refreshToken
	}
	authToken := RefreshToken(*authResult.RefreshToken)
	return &authToken
}

func (m *userManager) toUserSession(authResult *cognitoidentityprovider.AuthenticationResultType, refreshToken *RefreshToken) (UserSession, error) {
	if authResult == nil {
		return UserSession{}, errors.New("empty_auth")
	}
	if authResult.AccessToken == nil ||
		authResult.ExpiresIn == nil {
		return UserSession{}, errors.New("invalid_auth")
	}
	actualRefreshToken := m.extractRefreshToken(authResult, refreshToken)
	if actualRefreshToken == nil {
		return UserSession{}, errors.New("empty_refresh_token")
	}

	return UserSession{
		AccessToken:  AccessToken(*authResult.AccessToken),
		RefreshToken: *actualRefreshToken,
		ExpiresIn:    *authResult.ExpiresIn,
	}, nil
}
