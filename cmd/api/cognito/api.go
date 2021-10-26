package cognito

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/pkg/errors"
)

type UserManager interface {
	adminOperations
	userOperations
}

type Options struct {
	ClientID, PoolID, Region, AccessKey, SecretKey string
}

type userManager struct {
	appClientID   *string
	poolID        *string
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
}

func NewUserManager(opt Options) (UserManager, error) {
	conf := &aws.Config{
		Region:      aws.String(opt.Region),
		Credentials: credentials.NewStaticCredentials(opt.AccessKey, opt.SecretKey, ""),
	}

	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, errors.Wrap(err, "create_aws_session")
	}

	cognitoClient := cognitoidentityprovider.New(sess)

	return &userManager{
		appClientID:   &opt.ClientID,
		poolID:        &opt.PoolID,
		cognitoClient: cognitoClient,
	}, nil
}
