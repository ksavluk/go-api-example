package cognito

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/pkg/errors"
)

type adminOperations interface {
	GetUserUUIDByEmail(email string) (*string, error)
}

func (m *userManager) GetUserUUIDByEmail(email string) (*string, error) {
	filter := fmt.Sprintf("email = \"%s\"", email)
	searchOutput, err := m.cognitoClient.ListUsers(&cognitoidentityprovider.ListUsersInput{
		AttributesToGet: nil,
		Filter:          &filter,
		Limit:           nil,
		PaginationToken: nil,
		UserPoolId:      m.poolID,
	})
	if err != nil {
		return nil, err
	}

	users := searchOutput.Users
	if len(users) == 0 {
		return nil, nil
	}
	if len(users) == 1 {
		user := users[0]
		return user.Username, nil
	}
	return nil, errors.New("unable_to_find_user")
}
